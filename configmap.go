package k8stools
import (
    metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
    "k8s.io/client-go/kubernetes"
    "context"
    coreV1 "k8s.io/api/core/v1"
)

func UpdateConfigMap(c *kubernetes.Clientset,ns,configmap string,kv map[string]string) error {
    cm,err := ConfigMapInfo(c,ns,configmap)
    if err != nil {
        return err
    }
    cm.Data = kv
    _,err = c.CoreV1().ConfigMaps(ns).Update(context.TODO(),cm,metaV1.UpdateOptions{})
    if err !=nil {
        return err
    }
    return nil
}
func ConfigMapList(c *kubernetes.Clientset,ns string) (*coreV1.ConfigMapList,error) {
    var (
        configmaps *coreV1.ConfigMapList = &coreV1.ConfigMapList{}
        err error
    )
    if ns != "" {
        configmaps,err = c.CoreV1().ConfigMaps(ns).List(context.Background(),metaV1.ListOptions{})
    } else {
        configmaps,err = c.CoreV1().ConfigMaps("").List(context.Background(),metaV1.ListOptions{})
    }
    if err != nil {
        return nil,err
    }
    return  configmaps,nil
}
func ConfigMapInfo(c *kubernetes.Clientset,ns,mapname string) (*coreV1.ConfigMap,error) {
    mapinfo,err := c.CoreV1().ConfigMaps(ns).Get(context.Background(),mapname,metaV1.GetOptions{})
    if err != nil {
        return nil,err
    }
    return mapinfo,nil
}
func CreateConfigMap(c *kubernetes.Clientset,ns,configmap string,kv map[string]string) error {
    configMap := &coreV1.ConfigMap{
        ObjectMeta: metaV1.ObjectMeta{
            Name: configmap,
        },
        Data: kv,
    }
    _,err := c.CoreV1().ConfigMaps(ns).Create(context.TODO(),configMap,metaV1.CreateOptions{})
    if err != nil {
        return err
    }
    return nil
}
