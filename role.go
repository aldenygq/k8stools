package k8stools
import (
    metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
    rbacV1 "k8s.io/api/rbac/v1"
    "context"
    "k8s.io/client-go/kubernetes"
)
func RoleBindingInfo(c *kubernetes.Clientset,ns,name string) (*rbacV1.RoleBinding,error) {
    rolebinding,err := c.RbacV1().RoleBindings(ns).Get(context.TODO(),name,metaV1.GetOptions{})
    if err != nil {
        return nil,err
    }
    return rolebinding,nil
}
func RoleList(c *kubernetes.Clientset,ns string) (*rbacV1.RoleList,error) {
    roles,err := c.RbacV1().Roles(ns).List(context.TODO(),metaV1.ListOptions{})
    if err != nil {
        return nil,err
    }
    return roles,nil
}

func RoleBindingList(c *kubernetes.Clientset,ns string) (*rbacV1.RoleBindingList,error) {
    rolebindings,err := c.RbacV1().RoleBindings(ns).List(context.TODO(),metaV1.ListOptions{})
    if err != nil {
        return nil,err
    }
    return rolebindings,nil
}

func DeleteRoleBinding(c *kubernetes.Clientset,ns,name string) error {
    err := c.RbacV1().RoleBindings(ns).Delete(context.TODO(),name,metaV1.DeleteOptions{})
    if err != nil {
        return err
    }
    return nil
}
func DeleteRole(c *kubernetes.Clientset,ns,name string) error {
    err := c.RbacV1().Roles(ns).Delete(context.TODO(),name,metaV1.DeleteOptions{})
    if err != nil {
        return err
    }

    return nil
}
func ClusterRoleList(c *kubernetes.Clientset) (*rbacV1.ClusterRoleList,error) {
    clusterroles,err := c.RbacV1().ClusterRoles().List(context.TODO(),metaV1.ListOptions{})
    if err != nil {
        return nil,err
    }

    return clusterroles,nil
}
func ClusterRoleBindingList(c *kubernetes.Clientset) (*rbacV1.ClusterRoleBindingList,error) {
    clusterrolebindings,err := c.RbacV1().ClusterRoleBindings().List(context.TODO(),metaV1.ListOptions{})
    if err != nil {
        return nil,err
    }

    return clusterrolebindings,nil
}
func ClusterRoleBindingInfo(c *kubernetes.Clientset,name string) (*rbacV1.ClusterRoleBinding,error) {
    clusterrolebinding,err := c.RbacV1().ClusterRoleBindings().Get(context.TODO(),name,metaV1.GetOptions{})
     if err != nil {
         return nil,err
     }

     return clusterrolebinding,nil
}

func DeleteClusterRole(c *kubernetes.Clientset,name string) error {
    err := c.RbacV1().ClusterRoles().Delete(context.TODO(),name,metaV1.DeleteOptions{})
    if err != nil {
        return err
    }

    return nil
}

func DeleteClusterRoleBinding(c *kubernetes.Clientset,name string) error {
    err := c.RbacV1().ClusterRoleBindings().Delete(context.TODO(),name,metaV1.DeleteOptions{})
    if err != nil {
        return err
    }
    return nil
}

func RoleInfo(c *kubernetes.Clientset,ns,role string) (*rbacV1.Role,error) {
    roleinfo,err := c.RbacV1().Roles(ns).Get(context.TODO(),role,metaV1.GetOptions{})
    if err != nil {
         return nil,err
    }
    return roleinfo,nil
}
func ClusterRoleInfo(c *kubernetes.Clientset,rolename string) (*rbacV1.ClusterRole,error) {
    clusterrole,err := c.RbacV1().ClusterRoles().Get(context.TODO(),rolename,metaV1.GetOptions{})
    if err != nil {
        return nil,err
    }
    return clusterrole,nil
}
