package k8stools
import (
    "k8s.io/client-go/kubernetes"
    metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
    "context"
    coreV1 "k8s.io/api/core/v1"
)
func DeletePod(c *kubernetes.Clientset,ns,podname string) error {
    err := c.CoreV1().Pods(ns).Delete(context.TODO(), podname, metaV1.DeleteOptions{})
    if err != nil {
        return err
    }
    return nil
}
func PodList(c *kubernetes.Clientset,ns string) (*coreV1.PodList,error) {
    var (
        podlist *coreV1.PodList = &coreV1.PodList{}
        err error
    )
    // 列出Pods
    if ns != "" {
        podlist, err = c.CoreV1().Pods(ns).List(context.TODO(), metaV1.ListOptions{})
    } else {
        podlist, err = c.CoreV1().Pods("").List(context.TODO(), metaV1.ListOptions{})
    }
    if err != nil {
        return nil,err
    }
    return podlist,nil
}


func PodInfo(c *kubernetes.Clientset,ns,podname string) (*coreV1.Pod,error) {
    pod, err := c.CoreV1().Pods(ns).Get(context.TODO(), podname, metaV1.GetOptions{})
    if err != nil {
        return nil,err
    }
    return pod,nil
}
