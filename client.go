package k8stools
import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"time"

	appsV1 "k8s.io/api/apps/v1"
	batchV1 "k8s.io/api/batch/v1"
	coreV1 "k8s.io/api/core/v1"
	networkV1 "k8s.io/api/networking/v1"
	storageV1 "k8s.io/api/storage/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	rbacV1 "k8s.io/api/rbac/v1"
	//rbacV1beta1 "k8s.io/api/rbac/v1beta1"
)
var K8sClient *kubernetes.Clientset 
func NewK8sClient(kubeconfig string) (*kubernetes.Clientset,error){
  var err error 
	// 使用kubeconfig文件来获取客户端配置
	config, err := clientcmd.RESTConfigFromKubeConfig([]byte(kubeconfig))
	if err != nil {
		return nil,err 
	}
 
	// 根据客户端配置创建一个Kubernetes客户端
	K8sClient, err = kubernetes.NewForConfig(config)
	if err != nil {
		return nil,err 
	}
 
	return K8sClient,nil 
}
