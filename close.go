package k8stools
func (c *K8sClient) CloseClient() {
     c.Client = nil
}
