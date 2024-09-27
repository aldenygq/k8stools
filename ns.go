package k8stools
import (
    coreV1 "k8s.io/api/core/v1"
    "context"
    metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)
func (c *K8sClient) CreateNs(ns string) error {
    //defer c.CloseClient()
    // 创建namespace的spec
	var namespace coreV1.Namespace
	namespace.Name = ns
	_, err := c.Client.CoreV1().Namespaces().Create(context.TODO(),&namespace,metaV1.CreateOptions{})
    if err != nil {
        return err
    }
	return nil
}
func (c *K8sClient) NsInfo(ns string) (*coreV1.Namespace,error) {
     //defer c.CloseClient()
     namespaceInfo,err := c.Client.CoreV1().Namespaces().Get(context.TODO(),ns,metaV1.GetOptions{})
     if err != nil {
         return namespaceInfo,err
     }
     return namespaceInfo,nil
}

func (c *K8sClient) NsInfoList() (*coreV1.NamespaceList,error) {
    //defer c.CloseClient()
    namespaceList,err := c.Client.CoreV1().Namespaces().List(context.TODO(), metaV1.ListOptions{})
    if err != nil {
        return namespaceList,err
    }
    return namespaceList,nil
}

func (c *K8sClient) NsList() ([]string,error) {
    var nslist []string = make([]string,0)
    nsinfos,err := c.NsInfoList()
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

func (c *K8sClient) DeleteNamespace(ns string) error {
    err := c.Client.CoreV1().Namespaces().Delete(context.TODO(), ns, metaV1.DeleteOptions{})
    if err != nil {
        return err
    }

    return nil
}
