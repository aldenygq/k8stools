package k8stools
import (
    "k8s.io/client-go/tools/clientcmd"
    "k8s.io/client-go/kubernetes"
)

//集群节点列表
func K8sVersion(c *kubernetes.Clientset) (string,error) {
     // 获取版本
     serverVersion, err := c.Discovery().ServerVersion()
     if err != nil {
         return "",err
     }

    return serverVersion.String(),nil
}
