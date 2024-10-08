package k8stools
import (
    "fmt"
    "testing"
    "encoding/json"
)
//go test -v -test.run Test_NodeInfoList
func Test_NodeInfoList(t *testing.T) {
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

    nodelist,err := NodeInfoList(client)
    if err != nil {
        fmt.Printf("get nodes failed failed:%v\n",err)
        return
    }

    fmt.Printf("node list:%v\n",nodelist)
}

//go test -v -test.run Test_NodeAddrInfos
func Test_NodeAddrInfos(t *testing.T) {
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
    nodeaddrinfos,err := NodeAddrInfos(client)
    if err != nil {
        fmt.Printf("get node addr info failed:%v\n",err)
        return
    }
    addrinfos,_ := json.Marshal(nodeaddrinfos)
    fmt.Printf("node list:%v\n",string(addrinfos))
}
//go test -v -test.run Test_NodeInfo
func Test_NodeInfo(t *testing.T) {
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

    nodeinfo,err := NodeInfo(client,"ip-10-255-143-12.ap-southeast-1.compute.internal")
    if err != nil {
        fmt.Printf("get node info failed:%v\n",err)
        return
    }

    fmt.Printf("node info:%v\n",nodeinfo)
}
