func CreateNs(client *ns string) error {
	// 创建namespace的spec
	var namespace coreV1.Namespace
	namespace.Name = ns
	defer k.CloseClient()
	_, err := k.Client.CoreV1().Namespaces().Create(context.TODO(),&namespace,metaV1.CreateOptions{})
    if err != nil {
        return err 
    }
	return nil 
}
