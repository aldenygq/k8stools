package k8stools
import (
    appsV1 "k8s.io/api/apps/v1"
    batchV1 "k8s.io/api/batch/v1"
    metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
    "k8s.io/client-go/kubernetes"
    "context"
    "k8s.io/apimachinery/pkg/util/intstr"
    "k8s.io/apimachinery/pkg/types"
    "fmt"
    "time"
    "github.com/aldenygq/toolkits"
)


func DeploymentList(c *kubernetes.Clientset,ns string) (*appsV1.DeploymentList,error) {
    var (
        deployments *appsV1.DeploymentList = &appsV1.DeploymentList{}
        err error
    )

    if ns != "" {
        deployments,err = c.AppsV1().Deployments(ns).List(context.Background(),metaV1.ListOptions{})
    } else {
        deployments,err = c.AppsV1().Deployments("").List(context.Background(),metaV1.ListOptions{})
    }
    if err != nil {
        return nil,err
    }
    return  deployments,nil
}

func StatefulSetList(c *kubernetes.Clientset,ns string) (*appsV1.StatefulSetList,error) {
    var (
        statefulsets *appsV1.StatefulSetList = &appsV1.StatefulSetList{}
        err error
    )
    if ns != "" {
        statefulsets,err = c.AppsV1().StatefulSets(ns).List(context.Background(),metaV1.ListOptions{})
    } else {
        statefulsets,err = c.AppsV1().StatefulSets("").List(context.Background(),metaV1.ListOptions{})
    }
    if err != nil {
        return nil,err
    }
    return  statefulsets,nil
}

func JobList(c *kubernetes.Clientset,ns string) (*batchV1.JobList,error) {
    var (
        jobs *batchV1.JobList = &batchV1.JobList{}
        err error
    )
    if ns != "" {
        jobs,err = c.BatchV1().Jobs(ns).List(context.Background(),metaV1.ListOptions{})
    } else {
        jobs,err = c.BatchV1().Jobs("").List(context.Background(),metaV1.ListOptions{})
    }
    if err != nil {
        return nil,err
    }
    return  jobs,nil
}

func CronJobList(c *kubernetes.Clientset,ns string) (*batchV1.CronJobList,error) {
    var (
        cronjobs *batchV1.CronJobList = &batchV1.CronJobList{}
        err error
    )
    if ns != "" {
        cronjobs,err = c.BatchV1().CronJobs(ns).List(context.Background(),metaV1.ListOptions{})
    } else {
        cronjobs,err = c.BatchV1().CronJobs("").List(context.Background(),metaV1.ListOptions{})
    }
    if err != nil {
        return nil,err
    }
    return  cronjobs,nil
}

func DeploymentInfo(c *kubernetes.Clientset,ns,deployname string) (*appsV1.Deployment,error) {
    deployinfo,err := c.AppsV1().Deployments(ns).Get(context.Background(), deployname, metaV1.GetOptions{})
    if err != nil {
        return nil,err
    }
    return deployinfo,nil
}
func StatefulSetInfo(c *kubernetes.Clientset,ns,statefulset string) (*appsV1.StatefulSet,error) {
    statefulsetinfo,err := c.AppsV1().StatefulSets(ns).Get(context.Background(), statefulset, metaV1.GetOptions{})
    if err != nil {
        return nil,err
    }
    return statefulsetinfo,nil
}
func DaemonSetInfo(c *kubernetes.Clientset,ns,daemonset string) (*appsV1.DaemonSet,error) {
    daemonsetinfo,err := c.AppsV1().DaemonSets(ns).Get(context.Background(),daemonset,metaV1.GetOptions{})
    if err != nil {
        return nil,err
    }
    return daemonsetinfo,nil
}
func JobInfo(c *kubernetes.Clientset,ns,job string) (*batchV1.Job,error) {
    jobinfo,err := c.BatchV1().Jobs(ns).Get(context.Background(),job,metaV1.GetOptions{})
    if err != nil {
        return nil,err
    }
    return jobinfo,nil
}
func CronJobInfo(c *kubernetes.Clientset,ns,cronjob string) (*batchV1.CronJob,error) {
    cronjobinfo,err := c.BatchV1().CronJobs(ns).Get(context.Background(),cronjob,metaV1.GetOptions{})
    if err != nil {
        return nil,err
    }
    return cronjobinfo,nil
}

func DeleteDeployment(c *kubernetes.Clientset,ns,name string) error {
    err := c.AppsV1().Deployments(ns).Delete(context.TODO(), name, metaV1.DeleteOptions{GracePeriodSeconds: toolkits.Int64ToPointInt64(0)})
    if err != nil {
        return err
    }
    return nil
}
func DeleteStatefulSet(c *kubernetes.Clientset,ns,name string) error {
    err := c.AppsV1().StatefulSets(ns).Delete(context.TODO(), name, metaV1.DeleteOptions{GracePeriodSeconds: toolkits.Int64ToPointInt64(0)})
    if err != nil {
        return err
    }
    return nil
}
func DeleteDaemonSet(c *kubernetes.Clientset,ns,name string) error {
    err := c.AppsV1().DaemonSets(ns).Delete(context.TODO(), name, metaV1.DeleteOptions{GracePeriodSeconds: toolkits.Int64ToPointInt64(0)})
    if err != nil {
        return err
    }
    return nil
}

func DeleteJob(c *kubernetes.Clientset,ns,name string) error {
    err := c.BatchV1().Jobs(ns).Delete(context.TODO(), name, metaV1.DeleteOptions{GracePeriodSeconds: toolkits.Int64ToPointInt64(0)})
    if err != nil {
        return err
    }
    return nil
}
func DeleteCronJob(c *kubernetes.Clientset,ns,name string) error {
    err := c.BatchV1().CronJobs(ns).Delete(context.TODO(), name, metaV1.DeleteOptions{GracePeriodSeconds: toolkits.Int64ToPointInt64(0)})
    if err != nil {
        return err
    }
    return nil
}

func DeployRollUpdate(c *kubernetes.Clientset,ns,deployment string) error {
    patch := fmt.Sprintf(`{"spec":{"template":{"metadata":{"annotations":{"date": "%s"}}}}}`, time.Now().String())
    _, err := c.AppsV1().Deployments(ns).Patch(context.TODO(), deployment, types.StrategicMergePatchType, []byte(patch), metaV1.PatchOptions{})
    if err != nil {
        return err
    }
    return nil
}

func DaemonSetRollUpdate(c *kubernetes.Clientset,daemonset,ns string) error {
    dst,err := DaemonSetInfo(c,ns,daemonset)
    if err != nil {
        return err
    }
    dst.Spec.UpdateStrategy = appsV1.DaemonSetUpdateStrategy{
        Type: appsV1.RollingUpdateDaemonSetStrategyType,
        RollingUpdate: &appsV1.RollingUpdateDaemonSet{
            MaxUnavailable: func(i intstr.IntOrString) *intstr.IntOrString { return &i }(intstr.FromInt(1)),
        },
    }
    _, err = c.AppsV1().DaemonSets(ns).Update(context.TODO(), dst, metaV1.UpdateOptions{})
    if err != nil {
        return err
    }
    return nil
}

func StatefulSetRollUpdate(c *kubernetes.Clientset,statefulset,ns string) error {
    st,err := StatefulSetInfo(c,ns,statefulset)
    if err != nil {
        return err
    }
    st.Spec.Template.Labels["timestamp"] = fmt.Sprint(time.Now().Unix())
    _, err = c.AppsV1().StatefulSets(ns).Update(context.TODO(), st, metaV1.UpdateOptions{})
    if err != nil {
        return err
    }
    return nil
}
