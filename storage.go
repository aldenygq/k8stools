package k8stools
import (
    "k8s.io/client-go/kubernetes"
    coreV1 "k8s.io/api/core/v1"
    "context"
    metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
    "github.com/aldenygq/toolkits"
    storageV1 "k8s.io/api/storage/v1"
)

func PvInfo(c *kubernetes.Clientset,pvname string) (*coreV1.PersistentVolume,error) {
    pvinfo,err := c.CoreV1().PersistentVolumes().Get(context.Background(),pvname,metaV1.GetOptions{})
    if err != nil {
        return nil,err
    }
    return pvinfo,nil
}
func StorageClassInfo(c *kubernetes.Clientset,ns,name string) (*storageV1.StorageClass,error) {
    storageClass, err := c.StorageV1().StorageClasses().Get(context.TODO(), name, metaV1.GetOptions{})
    if err != nil {
        return nil,err
    }
    return storageClass,nil
}

func PvcInfo(c *kubernetes.Clientset,ns,pvcname string) (*coreV1.PersistentVolumeClaim,error) {
    pvcinfo,err := c.CoreV1().PersistentVolumeClaims(ns).Get(context.Background(),pvcname,metaV1.GetOptions{})
    if err != nil {
        return nil,err
    }
    return pvcinfo,nil
}
func DeletePv(c *kubernetes.Clientset,name string) error {
    err := c.CoreV1().PersistentVolumes().Delete(context.TODO(), name, metaV1.DeleteOptions{GracePeriodSeconds: toolkits.Int64ToPointInt64(0)})
    if err != nil {
        return err
    }
    return nil
}

func DeleteStorageClass(c *kubernetes.Clientset,name string) error {
    err := c.StorageV1().StorageClasses().Delete(context.TODO(), name, metaV1.DeleteOptions{GracePeriodSeconds: toolkits.Int64ToPointInt64(0)})
    if err != nil {
        return err
    }
    return nil
}

func DeletePvc(c *kubernetes.Clientset,ns,name string) error {
    err := c.CoreV1().PersistentVolumeClaims(ns).Delete(context.TODO(), name, metaV1.DeleteOptions{GracePeriodSeconds: toolkits.Int64ToPointInt64(0)})
    if err != nil {
        return err
    }
    return nil
}

func StorageClassList(c *kubernetes.Clientset) (*storageV1.StorageClassList,error) {
    var (
        storageclasses *storageV1.StorageClassList = &storageV1.StorageClassList{}
        err error
    )
    storageclasses,err = c.StorageV1().StorageClasses().List(context.Background(),metaV1.ListOptions{})
    if err != nil {
        return nil,err
    }
    return  storageclasses,nil
}

func PvList(c *kubernetes.Clientset) (*coreV1.PersistentVolumeList,error) {
    var (
        pvs *coreV1.PersistentVolumeList = &coreV1.PersistentVolumeList{}
        err error
    )
    pvs,err = c.CoreV1().PersistentVolumes().List(context.Background(),metaV1.ListOptions{})
    if err != nil {
        return nil,err
    }
    return  pvs,nil
}

func PvcList(c *kubernetes.Clientset,ns string) (*coreV1.PersistentVolumeClaimList,error) {
    var (
        pvcs *coreV1.PersistentVolumeClaimList = &coreV1.PersistentVolumeClaimList{}
        err error
    )
    if ns != "" {
        pvcs,err = c.CoreV1().PersistentVolumeClaims(ns).List(context.Background(),metaV1.ListOptions{})
    } else {
        pvcs,err = c.CoreV1().PersistentVolumeClaims("").List(context.Background(),metaV1.ListOptions{})
    }
    if err != nil {
        return nil,err
    }
    return  pvcs,nil
}
