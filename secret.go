package k8stools
import (
        "k8s.io/client-go/kubernetes"
        coreV1 "k8s.io/api/core/v1"
        "context"
        metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
        "fmt"
        "github.com/aldenygq/toolkits"
)
func UpdateSecretByOpaque(c *kubernetes.Clientset,ns,secretname string,kv map[string][]byte) error {
    secret := &coreV1.Secret{
        ObjectMeta: metaV1.ObjectMeta{
            Name: secretname,
        },
        Type: "Opaque",
        Data: kv,
    }
    _,err := c.CoreV1().Secrets(ns).Update(context.TODO(), secret, metaV1.UpdateOptions{})
    if err != nil {
        return err
    }
    return nil
}
func CreateSecretByOpaque(c *kubernetes.Clientset,ns,secretname string,kv map[string][]byte) error {
    secret := &coreV1.Secret{
        ObjectMeta: metaV1.ObjectMeta{
            Name: secretname,
        },
        Type: "Opaque",
        Data: kv,
    }
    _,err := c.CoreV1().Secrets(ns).Create(context.TODO(), secret, metaV1.CreateOptions{})
    if err != nil {
        return err
    }
    return nil
}

func UpdateSecretByTlsCert(c *kubernetes.Clientset,ns,secretname,cert,key string) error {
    data := map[string][]byte{
        coreV1.TLSCertKey:       []byte(cert),
        coreV1.TLSPrivateKeyKey: []byte(key),
    }
    secret := &coreV1.Secret{
        ObjectMeta: metaV1.ObjectMeta{
            Name:      secretname,
            Namespace: ns,
        },
        Type: coreV1.SecretTypeTLS,
        Data: data,
    }
    _, err := c.CoreV1().Secrets(ns).Update(context.TODO(),secret,metaV1.UpdateOptions{})
    if err != nil {
        return err
    }
    return nil
}

func CreateSecretByTlsCert(c *kubernetes.Clientset,ns,secretname,cert,key string) error {
    data := map[string][]byte{
        coreV1.TLSCertKey:       []byte(cert),
        coreV1.TLSPrivateKeyKey: []byte(key),
    }
    secret := &coreV1.Secret{
        ObjectMeta: metaV1.ObjectMeta{
            Name:      secretname,
            Namespace: ns,
        },
        Type: coreV1.SecretTypeTLS,
        Data: data,
    }
    _, err := c.CoreV1().Secrets(ns).Create(context.TODO(),secret,metaV1.CreateOptions{})
    if err != nil {
        return err
    }
    return nil
}
func UpdateSecretByImageCert(c *kubernetes.Clientset,ns,secretname,url,user,password string) error {
    dockerConfigJson := fmt.Sprintf(`{"auths":{"%v":{"username":"%v","password":"%v"}}}"`,url,user,password)
    secret := &coreV1.Secret{
        ObjectMeta: metaV1.ObjectMeta{
            Name: secretname,
        },
        Type: coreV1.SecretTypeDockerConfigJson,
        Data: map[string][]byte{
            ".dockerconfigjson": []byte(dockerConfigJson),
        },
    }
    _, err := c.CoreV1().Secrets(ns).Update(context.TODO(), secret, metaV1.UpdateOptions{})
    if err != nil {
        return err
    }

    return nil
}

func CreateSecretByImageCert(c *kubernetes.Clientset,ns,secretname,url,user,password string) error {
     dockerConfigJson := fmt.Sprintf(`{"auths":{"%v":{"username":"%v","password":"%v"}}}"`,url,user,password)
     secret := &coreV1.Secret{
         ObjectMeta: metaV1.ObjectMeta{
             Name: secretname,
         },
         Type: coreV1.SecretTypeDockerConfigJson,
         Data: map[string][]byte{
             ".dockerconfigjson": []byte(dockerConfigJson),
         },
     }

     _, err := c.CoreV1().Secrets(ns).Create(context.TODO(), secret, metaV1.CreateOptions{})
     if err != nil {
         return err
     }

     return nil
}
func SecretInfo(c *kubernetes.Clientset,ns,sectet string) (*coreV1.Secret,error) {
    secretinfo,err := c.CoreV1().Secrets(ns).Get(context.Background(),sectet,metaV1.GetOptions{})
    if err != nil {
        return nil,err
    }
    return secretinfo,nil
}

func SecretList(c *kubernetes.Clientset,ns string) (*coreV1.SecretList,error) {
    var (
        secrets *coreV1.SecretList = &coreV1.SecretList{}
        err error
    )
    if ns != "" {
        secrets,err = c.CoreV1().Secrets(ns).List(context.Background(),metaV1.ListOptions{})
    } else {
        secrets,err = c.CoreV1().Secrets("").List(context.Background(),metaV1.ListOptions{})
    }
    if err != nil {
        return nil,err
    }
    return  secrets,nil
}
func DeleteSecret(c *kubernetes.Clientset,ns,name string) error {
    err := c.CoreV1().Secrets(ns).Delete(context.TODO(), name, metaV1.DeleteOptions{GracePeriodSeconds: toolkits.Int64ToPointInt64(0)})
    if err != nil {
        return err
    }
    return nil
}
