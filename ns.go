package k8stools
import (
    coreV1 "k8s.io/api/core/v1"
    "context"
    metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
    "k8s.io/client-go/kubernetes"
)
func CreateNs(c *kubernetes.Clientset,ns string) error {
    // 创建namespace的spec
	var namespace coreV1.Namespace
	namespace.Name = ns
	_, err := c.CoreV1().Namespaces().Create(context.TODO(),&namespace,metaV1.CreateOptions{})
    if err != nil {
        return err
    }
	return nil
}
func NsInfo(c *kubernetes.Clientset,ns string) (*coreV1.Namespace,error) {
     namespaceInfo,err := c.CoreV1().Namespaces().Get(context.TODO(),ns,metaV1.GetOptions{})
     if err != nil {
         return nil,err
     }
     return namespaceInfo,nil
}

func NsInfoList(c *kubernetes.Clientset) (*coreV1.NamespaceList,error) {
    //defer c.CloseClient()
    namespaceList,err := c.CoreV1().Namespaces().List(context.TODO(), metaV1.ListOptions{})
    if err != nil {
        return nil,err
    }
    return namespaceList,nil
}

func NsList(c *kubernetes.Clientset) ([]string,error) {
    var nslist []string = make([]string,0)
    nsinfos,err := NsInfoList(c)
    if err != nil {
        return nil,err
    }
    if len(nsinfos.Items) <= 0 {
        return nslist,nil
    }
    for _,v := range nsinfos.Items {
        nslist = append(nslist,v.Name)
    }

    return nslist,nil
}

func DeleteNamespace(c *kubernetes.Clientset,ns string) error {
    err := c.CoreV1().Namespaces().Delete(context.TODO(), ns, metaV1.DeleteOptions{})
    if err != nil {
        return err
    }

    return nil
}
