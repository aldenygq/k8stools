package k8stools
import (
    "k8s.io/client-go/kubernetes"
    coreV1 "k8s.io/api/core/v1"
    metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
    "context"
)
func EventInfo(c *kubernetes.Clientset,ns,filter string) (*coreV1.EventList,error) {
    var (
        watcher *coreV1.EventList = &coreV1.EventList{}
        err error
    )
    if ns == "" {
        // 获取事件的watcher
        watcher, err = c.CoreV1().Events("").List(context.Background(), metaV1.ListOptions{})
    } else {
        watcher, err = c.CoreV1().Events(ns).List(context.TODO(), metaV1.ListOptions{
            FieldSelector: filter,
        })
    }
    if err != nil {
        return nil,err
    }
    return watcher,nil
}
