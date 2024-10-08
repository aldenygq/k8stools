package k8stools
import (
    "fmt"
    "testing"
)
//go test -v -test.run Test_NewClient
func Test_NewClient(t *testing.T) {
    filepath := fmt.Sprintf("~/.kube/config")
    conf,err := GetKubeconfigByFile(filepath)
    if err != nil {
        fmt.Printf("get kubeconfig failed:%v\n",err)
        return
    }

    client,err := NewK8sClient(conf)
    if err != nil || client == nil {
        fmt.Printf("new client failed:%v\n",err)
        return
    }

    fmt.Printf("new client success\n")
}
