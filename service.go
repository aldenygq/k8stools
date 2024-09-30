package k8stools
import (
    networkV1 "k8s.io/api/networking/v1"
    "k8s.io/client-go/kubernetes"
    coreV1 "k8s.io/api/core/v1"
    "context"
    metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
    "github.com/aldenygq/toolkits"
)
func SvcList(c *kubernetes.Clientset,ns string) (*coreV1.ServiceList,error) {
    var (
        svcs *coreV1.ServiceList = &coreV1.ServiceList{}
        err error
    )
    if ns != "" {
        svcs,err = c.CoreV1().Services(ns).List(context.Background(),metaV1.ListOptions{})
    } else {
        svcs,err = c.CoreV1().Services("").List(context.Background(),metaV1.ListOptions{})
    }
    if err != nil {
        return nil,err
    }
    return  svcs,nil
}
func SvcInfo(c *kubernetes.Clientset,ns,svcname string)  (*coreV1.Service,error) {
    // 获取Service对象
    service, err := c.CoreV1().Services(ns).Get(context.Background(), svcname, metaV1.GetOptions{})
    if err != nil {
        return nil,err
    }
    return service,nil
}
func DeleteService(c *kubernetes.Clientset,ns,name string) error {
    err := c.CoreV1().Services(ns).Delete(context.TODO(), name, metaV1.DeleteOptions{GracePeriodSeconds: toolkits.Int64ToPointInt64(0)})
    if err != nil {
        return err
    }
    return nil
}
func DeleteIngress(c *kubernetes.Clientset,ns,name string) error {
    err := c.NetworkingV1().Ingresses(ns).Delete(context.TODO(), name, metaV1.DeleteOptions{GracePeriodSeconds:toolkits.Int64ToPointInt64(0)})
    if err != nil {
        return err
    }
    return nil
}
func IngressList(c *kubernetes.Clientset,ns string) (*networkV1.IngressList,error) {
    var (
        ingresses *networkV1.IngressList = &networkV1.IngressList{}
        err error
    )
    if ns != "" {
        ingresses,err = c.NetworkingV1().Ingresses(ns).List(context.Background(),metaV1.ListOptions{})
    } else {
        ingresses,err = c.NetworkingV1().Ingresses("").List(context.Background(),metaV1.ListOptions{})
    }
    if err != nil {
        return nil,err
    }
    return  ingresses,nil
}
func IngressInfo(c *kubernetes.Clientset,ns,ingressname string) (*networkV1.Ingress,error) {
    ingress, err := c.NetworkingV1().Ingresses(ns).Get(context.TODO(), ingressname, metaV1.GetOptions{})
    if err != nil {
        return nil,err
    }
    return ingress,nil
}
