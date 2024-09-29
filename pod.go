package k8stools
import (
    "k8s.io/client-go/kubernetes"
    metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
    "context"
)
func DeletePod(c *kubernetes.Clientset,ns,podname string) error {
    err := c.CoreV1().Pods(ns).Delete(context.TODO(), podname, metaV1.DeleteOptions{})
    if err != nil {
        return err
    }
    return nil
}
