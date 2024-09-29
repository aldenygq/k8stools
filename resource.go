package k8stools
import (
    appsV1 "k8s.io/api/apps/v1"
    batchV1 "k8s.io/api/batch/v1"
    coreV1 "k8s.io/api/core/v1"
    networkV1 "k8s.io/api/networking/v1"
    storageV1 "k8s.io/api/storage/v1"
    metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
    "k8s.io/client-go/kubernetes"
    "context"
    "errors"
    rbacV1 "k8s.io/api/rbac/v1"
)

func CreateResource(c *kubernetes.Clientset,obj interface{}) error {
    var err error
    switch obj.(type) {
    case coreV1.Namespace:
        _,err = c.CoreV1().Namespaces().Create(context.TODO(), obj.(*coreV1.Namespace),metaV1.CreateOptions{})
    case appsV1.Deployment:
        _,err = c.AppsV1().Deployments("").Create(context.TODO(), obj.(*appsV1.Deployment),metaV1.CreateOptions{})
    case appsV1.StatefulSet:
        _,err = c.AppsV1().StatefulSets("").Create(context.TODO(), obj.(*appsV1.StatefulSet),metaV1.CreateOptions{})
    case appsV1.DaemonSet:
        _,err = c.AppsV1().DaemonSets("").Create(context.TODO(), obj.(*appsV1.DaemonSet),metaV1.CreateOptions{})
    case batchV1.Job:
        _,err = c.BatchV1().Jobs("").Create(context.TODO(), obj.(*batchV1.Job),metaV1.CreateOptions{})
    case batchV1.CronJob:
        _,err = c.BatchV1().CronJobs("").Create(context.TODO(), obj.(*batchV1.CronJob),metaV1.CreateOptions{})
    case coreV1.Pod:
        _,err = c.CoreV1().Pods("").Create(context.TODO(), obj.(*coreV1.Pod),metaV1.CreateOptions{})
    case coreV1.Service:
        _,err = c.CoreV1().Services("").Create(context.TODO(), obj.(*coreV1.Service),metaV1.CreateOptions{})
    case networkV1.Ingress:
        _,err = c.NetworkingV1().Ingresses("").Create(context.TODO(), obj.(*networkV1.Ingress),metaV1.CreateOptions{})
    case coreV1.ConfigMap:
        _,err = c.CoreV1().ConfigMaps("").Create(context.TODO(), obj.(*coreV1.ConfigMap),metaV1.CreateOptions{})
    case coreV1.Secret:
        _,err = c.CoreV1().Secrets("").Create(context.TODO(), obj.(*coreV1.Secret),metaV1.CreateOptions{})
    case coreV1.PersistentVolumeClaim:
        _,err = c.CoreV1().PersistentVolumeClaims("").Create(context.TODO(), obj.(*coreV1.PersistentVolumeClaim),metaV1.CreateOptions{})
    case coreV1.PersistentVolume:
        _,err = c.CoreV1().PersistentVolumes().Create(context.TODO(), obj.(*coreV1.PersistentVolume),metaV1.CreateOptions{})
    case storageV1.StorageClass:
        _,err = c.StorageV1().StorageClasses().Create(context.TODO(), obj.(*storageV1.StorageClass),metaV1.CreateOptions{})
    case coreV1.ServiceAccount:
        _,err = c.CoreV1().ServiceAccounts("").Create(context.TODO(), obj.(*coreV1.ServiceAccount),metaV1.CreateOptions{})
    case rbacV1.ClusterRole:
        _,err = c.RbacV1().ClusterRoles().Create(context.TODO(), obj.(*rbacV1.ClusterRole), metaV1.CreateOptions{})
    default:
        return errors.New("resource type invaild")
    }
    if err != nil {
        return err
    }
    return nil
}
