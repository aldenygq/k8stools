package k8stools
import (
    "k8s.io/client-go/tools/clientcmd"
    "k8s.io/client-go/kubernetes"
    restclient "k8s.io/client-go/rest"
)
func NewK8sClient(conf *restclient.Config) (*kubernetes.Clientset,error){
	// 根据客户端配置创建一个Kubernetes客户端
	client, err := kubernetes.NewForConfig(conf)
	if err != nil {
		return nil,err
	}
	return client,nil
}
func GetKubeconfigByFile(filepath string) (*restclient.Config,error) {
    config, err := clientcmd.BuildConfigFromFlags("", filepath)
    if err != nil {
        return nil,err
    }
    return config,nil
}
func GetKubeconfigByKubeconfig(kubeconfig string) (*restclient.Config,error) {
    // 使用kubeconfig文件来获取客户端配置
    config, err := clientcmd.RESTConfigFromKubeConfig([]byte(kubeconfig))
    if err != nil {
        return nil,err
    }
    return config,nil
}

func Close(c *kubernetes.Clientset) {
	c = nil
}
