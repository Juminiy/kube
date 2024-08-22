package k8s_api

import (
	"context"
	"errors"
	"fmt"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/client-go/util/retry"
	"strconv"
	"strings"

	_ "k8s.io/client-go/plugin/pkg/client/auth"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8scli "k8s.io/client-go/kubernetes"
	typedappsv1 "k8s.io/client-go/kubernetes/typed/apps/v1"

	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"

	"kube/util"
	"path/filepath"
)

// move it to yaml
const (
	internalImageRegistry            = "192.168.31.242:8662"   // no http or https
	internalImageProject             = "kubesphere-io-centos7" // namespace override
	containerRequestResourceRatio    = 0.5                     // request = ratio * limit
	ContainerLimitDiskCacheDefaultGi = 8                       // ephemeral 8 GiB
)

type (
	DeploymentConfig struct {
		Namespace          string
		MetaName           string
		SpecReplicas       int32
		SpecSelectorLabels map[string]string
		SpecTemplateLabels map[string]string
		ContainerName      string
		ContainerImage     string
		ContainerCommand   []string
		ContainerPort      corev1.ContainerPort
		ContainerResource  *ResourceDecl

		CallBack // what should do after CRUD kube api

		k8s *k8scli.Clientset
		cli typedappsv1.DeploymentInterface
		app *appsv1.Deployment

		ctx context.Context // context in k8e
	}

	// Pod Config Decl
	EnvironmentDecl struct {
	}

	// Container Resource Limit
	ResourceDecl struct {
		CPU       float64 // VCPU Logical 					 /Core
		GPU       float64 // VGPU 							 /Core
		GMem      int64   // VGPU Self VMemory				 /Byte
		Mem       int64   // VMemory 				  			 /Byte
		DiskMount int64   // Volume (Minio Cluster s3fs) mount /Byte
		DiskCache int64   // Ephemeral mount 					 /Byte

		BindGPU   bool
		BindMount bool
	}

	CallBack struct {
		latest any
		err    error
		Create util.Func
		Update util.Func
		Delete util.Func
		List   util.Func
	}
)

func NewDeployment(c *DeploymentConfig) error {
	if validErr := c.validate(); validErr != nil {
		return validErr
	}

	restConfig, err := clientcmd.BuildConfigFromFlags(
		"",
		filepath.Join(homedir.HomeDir(), ".kube", "config"))
	if err != nil {
		return err
	}

	c.k8s, err = k8scli.NewForConfig(restConfig)
	if err != nil {
		return err
	}

	c.cli = c.k8s.AppsV1().Deployments(c.Namespace)

	c.app = &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: c.MetaName,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: util.NewInt32(c.SpecReplicas),
			Selector: &metav1.LabelSelector{
				MatchLabels: c.SpecSelectorLabels,
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: c.SpecTemplateLabels,
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:    c.ContainerName,
							Image:   c.ContainerImage,
							Command: c.ContainerCommand,
							Ports: []corev1.ContainerPort{
								c.ContainerPort,
							},
							Resources: corev1.ResourceRequirements{
								Limits: corev1.ResourceList{
									corev1.ResourceCPU:              resource.MustParse(getCPUStr(c.ContainerResource.CPU)),
									corev1.ResourceMemory:           resource.MustParse(getByteStr(c.ContainerResource.Mem)),
									corev1.ResourceEphemeralStorage: resource.MustParse(getByteStr(c.ContainerResource.DiskCache)),
								},
								Requests: corev1.ResourceList{
									corev1.ResourceCPU:              resource.MustParse(getCPUStr(c.ContainerResource.CPU / 2.0)),
									corev1.ResourceMemory:           resource.MustParse(getByteStr(c.ContainerResource.Mem / 2)),
									corev1.ResourceEphemeralStorage: resource.MustParse(getByteStr(c.ContainerResource.DiskCache / 2)),
								},
							},
						},
					},
				},
			},
		},
	}
	// GPU Relation
	if c.ContainerResource.BindGPU {
	}
	// Mount Relation
	if c.ContainerResource.BindMount {
		for i := range c.app.Spec.Template.Spec.Containers {
			c.app.Spec.Template.Spec.Containers[i].Resources.Limits[corev1.ResourceStorage] =
				resource.MustParse(getByteStr(c.ContainerResource.DiskMount))
			c.app.Spec.Template.Spec.Containers[i].Resources.Requests[corev1.ResourceStorage] =
				resource.MustParse(getByteStr(c.ContainerResource.DiskMount / 2))
		}
	}

	c.ctx = context.TODO()

	defaultCallBackFunc := func() error {
		if c.latest != nil {
			fmt.Println(c.latest)
		}
		c.latest = nil
		return nil
	}
	c.CallBack = CallBack{
		Create: defaultCallBackFunc,
		Update: defaultCallBackFunc,
		Delete: defaultCallBackFunc,
		List:   defaultCallBackFunc,
	}

	return nil
}

func (c *DeploymentConfig) validate() error {
	if c == nil {
		return errors.New("deployment config is null")
	}
	if c.ContainerResource == nil {
		return errors.New("container resource is null")
	}
	return nil
}

func (c *DeploymentConfig) Create() error {
	var createErr error
	c.latest, createErr = c.cli.Create(c.ctx, c.app, metav1.CreateOptions{})
	if createErr != nil {
		return createErr
	}

	return c.CallBack.Create()
}

func (c *DeploymentConfig) Update() error {
	retryErr := retry.RetryOnConflict(
		retry.DefaultRetry,
		func() error {
			var updateErr error
			c.latest, updateErr = c.cli.Update(c.ctx, c.app, metav1.UpdateOptions{})
			return updateErr
		},
	)
	if retryErr != nil {
		return retryErr
	}

	return c.CallBack.Update()
}

func (c *DeploymentConfig) Delete() error {
	deletePolicy := metav1.DeletePropagationForeground
	deleteErr := c.cli.Delete(c.ctx, c.MetaName, metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	})
	if deleteErr != nil {
		return deleteErr
	}

	return c.CallBack.Delete()
}

func (c *DeploymentConfig) List() error {
	var listErr error
	c.latest, listErr = c.cli.List(c.ctx, metav1.ListOptions{})
	if listErr != nil {
		return listErr
	}

	return c.CallBack.List()
}

// GetImageURL
// Example:
// 192.168.31.242:8662/kubesphere-io-centos7/haproxy:2.9.6-alpine
func GetImageURL(image string) string {
	return strings.Join(
		[]string{internalImageRegistry, internalImageProject, image},
		"/",
	)
}

// when k8s node deploy in physical machine, core is physical core
// when k8s node deploy in virtual machine, core is virtual core
func getCPUStr(coreNum float64) string {
	return strconv.FormatFloat(coreNum, 'f', 6, 64)
}

// B -> KiB/MiB/GiB/TiB/PiB/EiB
// currently not conv
func getByteStr(memByte int64) string {
	return strconv.FormatInt(memByte, 10)
}
