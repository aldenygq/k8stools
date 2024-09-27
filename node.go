package k8stools
import (
    coreV1 "k8s.io/api/core/v1"
    "context"
    metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
    "k8s.io/client-go/kubernetes"
    "errors"
    "github.com/aldenygq/toolkits"
)

//集群节点列表
func NodeInfoList(c *kubernetes.Clientset) (*coreV1.NodeList,error) {
    nodes, err := c.CoreV1().Nodes().List(context.TODO(), metaV1.ListOptions{})
    if err != nil {
        return nil,err
    }
    return nodes,nil
}
type NodeAddrInfo struct {
    NodeHostName []string `json:"node_name"`
    InternalIps []string `json:"internal_ips"`
    ExternalIps  []string `json:"external_ips"`
    InternalDNS []string `json:"internal_dns"`
    ExternalDNS []string `json:"external_dns"`
}
//节点地址信息，包含私网ip/公网ip/主机名/私网dns/公网dns
func NodeAddrInfos(c *kubernetes.Clientset) ([]*NodeAddrInfo,error) {
    var nodeaddrinfos []*NodeAddrInfo = make([]*NodeAddrInfo,0)
    nodes ,err := NodeInfoList(c)
    if err != nil {
        return nil,err
    }
    if len(nodes.Items) <= 0 {
        return nil,nil
    }
    for _,v := range nodes.Items {
        var (
            nodeaddr *NodeAddrInfo = &NodeAddrInfo{}
            internalips []string = make([]string,0)
            externalips []string = make([]string,0)
            hostnames []string = make([]string,0)
            internaldns []string = make([]string,0)
            externaldns []string = make([]string,0)
        )
        for _, addr := range v.Status.Addresses {
            switch addr.Type {
            case coreV1.NodeInternalIP:
                internalips = append(internalips,addr.Address)
            case coreV1.NodeExternalIP:
                externalips = append(externalips,addr.Address)
            case coreV1.NodeHostName:
                hostnames = append(hostnames,addr.Address)
            case coreV1.NodeInternalDNS:
                internaldns = append(internaldns,addr.Address)
            case coreV1.NodeExternalDNS:
                externaldns = append(externaldns,addr.Address)
            }
        }
        nodeaddr.InternalIps = internalips
        nodeaddr.ExternalIps = externalips
        nodeaddr.InternalDNS = internaldns
        nodeaddr.ExternalDNS = externaldns
        nodeaddr.NodeHostName = hostnames
        nodeaddrinfos = append(nodeaddrinfos,nodeaddr)
    }
    return nodeaddrinfos,nil
}

//节点信息
func NodeInfo(c *kubernetes.Clientset,nodename string) (*coreV1.Node,error) {
    node,err := c.CoreV1().Nodes().Get(context.TODO(),nodename,metaV1.GetOptions{})
    if err != nil {
        return nil,err
    }
    return node,nil
}
//更新节点信息
func PatchNodeInfo(c *kubernetes.Clientset,node *coreV1.Node) error {
   // 更新节点
   _, err := c.CoreV1().Nodes().Update(context.TODO(), node, metaV1.UpdateOptions{})
   if err != nil {
       return err
   }
   return nil
}
//更新节点标签
func PatchNodeLable(c *kubernetes.Clientset,nodename string,labels map[string]string) error {
    node,err := NodeInfo(c,nodename)
    if err != nil {
        return err
    }
   // 添加或更新标签
   if node.Labels == nil {
       node.Labels = map[string]string{}
   }
   for k,v := range labels {
       node.Labels[k] = v
   }
   // 更新节点
   err = PatchNodeInfo(c,node)
   if err != nil {
       return err
   }

   return nil
}

//更新节点污点
func PatchNodeTaint(c *kubernetes.Clientset,nodename string,taints map[string]string) error {
        node,err := NodeInfo(c,nodename)
     if err != nil {
         return err
     }
    for k,v := range taints {
        taint := &coreV1.Taint{
            Key:    k,
            Value:  v,
            Effect: coreV1.TaintEffectNoSchedule,
        }
        node.Spec.Taints = append(node.Spec.Taints, *taint)
    }
    // 更新节点
   err = PatchNodeInfo(c,node)
   if err != nil {
       return err
   }
    return nil
}
//节点调度
func PatchNodeSchedule(c *kubernetes.Clientset,nodename,schedulerule string) error {
    node,err := NodeInfo(c,nodename)
    if err != nil {
        return err
    }
    switch schedulerule {
    case "disable":
        node.Spec.Unschedulable = true
    case "enable":
        node.Spec.Unschedulable = false
    default:
        return errors.New("schedule rule invalid")
    }
   // 更新节点
   err = PatchNodeInfo(c,node)
   if err != nil {
       return err
   }
    return nil
}
//节点驱逐排水
func PatchNodeDrain(c *kubernetes.Clientset,nodename string) error {
    pods, err := c.CoreV1().Pods("").List(context.TODO(), metaV1.ListOptions{
        FieldSelector: "spec.nodeName=" + nodename,
    })
    if err != nil {
        return err
    }
    for _, pod := range pods.Items {
        err := DeletePod(c,pod.Namespace,pod.Name)
        if err != nil {
            return err
        }
    }
    return nil
}
//节点下pod列表
func PodsInNode(c *kubernetes.Clientset,nodename string) (*coreV1.PodList,error) {
    pods, err := c.CoreV1().Pods("").List(context.TODO(), metaV1.ListOptions{
        FieldSelector: "spec.nodeName=" + nodename,
    })
    if err != nil {
        return nil,err
    }
    return pods,nil
}
//删除节点
func DeleteNode(c *kubernetes.Clientset,nodename string) error {
    // 删除节点
    err := c.CoreV1().Nodes().Delete(context.TODO(), nodename, metaV1.DeleteOptions{
        GracePeriodSeconds: toolkits.Int64ToPointInt64(0), // 立即删除节点
    })
    if err != nil {
        return err
    }
    return nil
}
