package k8stools
import (
    "k8s.io/apimachinery/pkg/runtime"
    "k8s.io/client-go/kubernetes"
    coreV1 "k8s.io/api/core/v1"
    "context"
)

func PodLog(c *kubernetes.Clientset,ns,podname string) (runtime.Object,error) {
    podLogs,err := c.CoreV1().Pods(ns).GetLogs(podname, &coreV1.PodLogOptions{}).Do(context.TODO()).Get()
    if err != nil {
        return nil,err
    }
    return podLogs,nil
}
