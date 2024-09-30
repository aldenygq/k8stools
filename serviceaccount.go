package k8stools

import (
    coreV1 "k8s.io/api/core/v1"
    "k8s.io/client-go/kubernetes"
    metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
    "context"
)

func ServiceAccountList(c *kubernetes.Clientset) (*coreV1.ServiceAccountList,error) {
    serviceaccounts,err := c.CoreV1().ServiceAccounts("").List(context.TODO(),metaV1.ListOptions{})
    if err != nil {
        return nil,err
    }
    return serviceaccounts,nil
}
func DeleteServiceAccount(c *kubernetes.Clientset,ns,serviceaccount string) error {
    deletePolicy := metaV1.DeletePropagationForeground
    err := c.CoreV1().ServiceAccounts(ns).Delete(context.TODO(),serviceaccount,metaV1.DeleteOptions{
        PropagationPolicy: &deletePolicy,
    })
    if err != nil {
        return err
    }
    return nil
}
func ServiceAccountInfo(c *kubernetes.Clientset,ns,serviceaccount string) (*coreV1.ServiceAccount,error) {
    serviceaccountinfo,err := c.CoreV1().ServiceAccounts(ns).Get(context.Background(),serviceaccount,metaV1.GetOptions{})
    if err != nil {
        return nil,err
    }
    return serviceaccountinfo,nil
}
