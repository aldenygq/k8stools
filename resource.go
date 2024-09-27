package k8stools
import (

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
