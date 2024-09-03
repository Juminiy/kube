package pod_api

import (
	"context"
	"encoding/json"
	"errors"
	"path/filepath"
	"strconv"
	"strings"
	"sync"

	"github.com/Juminiy/kube/pkg/image_api/harbor_api"
	"github.com/Juminiy/kube/pkg/log_api/stdlog"
	"github.com/Juminiy/kube/pkg/util"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"k8s.io/client-go/util/retry"

	_ "k8s.io/client-go/plugin/pkg/client/auth"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8scli "k8s.io/client-go/kubernetes"
	typedappsv1 "k8s.io/client-go/kubernetes/typed/apps/v1"
)

// move it to yaml config
const (
	internalImageRegistry            = "192.168.31.242:8662" // no http or https
	containerRequestResourceRatio    = 0.5                   // request = ratio * limit
	ContainerLimitDiskCacheDefaultGi = 8                     // ephemeral 8 GiB
)

var (
	k8sClientSet struct {
		cs *k8scli.Clientset
		sync.Once
	}
)

type (
	DeploymentConfig struct {
		Namespace          string
		MetaName           string
		SpecReplicas       int32
		SpecSelectorLabels map[string]string
		SpecTemplateLabels map[string]string
		SpecHostNetwork    bool

		Container *ContainerConfig

		// do after kube api
		CallBack *CallBack `json:"CallBack,omitempty"`

		// global cli
		cliSet *k8scli.Clientset
		cli    typedappsv1.DeploymentInterface
		app    *appsv1.Deployment

		// context in k8e
		ctx context.Context
	}

	// Container Config Declaration in Pod
	ContainerConfig struct {
		Name            string
		Image           string
		Command         []string
		Args            string
		Ports           []corev1.ContainerPort
		Resource        *ResourceConfig
		SecurityContext *corev1.SecurityContext
		UserHost        *UserHostConfig
	}

	// Container Resource Declaration in Pod
	ResourceConfig struct {
		CPU       float64 // VCPU Logical 					 	/Core
		GPU       float64 // VGPU 							 	/Core
		GMem      int64   // VGPU Self VMemory				 	/Byte
		Mem       int64   // VMemory 				  			/Byte
		DiskMount int64   // Volume (Minio Cluster s3fs) mount 	/Byte
		DiskCache int64   // Ephemeral mount 					/Byte

		BindGPU   bool
		BindMount bool
	}

	// User Host Declaration in Pod
	UserHostConfig struct {
		HostName string
		UserName string
		Password string
	}

	CallBack struct {
		latest any
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

	k8sClientSet.Do(func() {
		restConfig, err := clientcmd.BuildConfigFromFlags(
			"",
			filepath.Join(homedir.HomeDir(), ".kube", "config"))
		util.SilentHandleError("init k8s client error", err)

		k8sClientSet.cs, err = k8scli.NewForConfig(restConfig)
		util.SilentHandleError("init k8s client error", err)
	})
	c.cliSet = k8sClientSet.cs

	c.cli = c.cliSet.AppsV1().Deployments(c.Namespace)

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
					Containers:  []corev1.Container{*NewContainer(c.Container)},
					HostNetwork: c.SpecHostNetwork,
				},
			},
		},
	}

	c.ctx = context.TODO()

	defaultCallBackFunc := func() error {
		if c.CallBack.latest != nil {
			stdlog.Info(c.CallBack.latest)
		}
		c.CallBack.latest = nil
		return nil
	}
	c.CallBack = &CallBack{
		Create: defaultCallBackFunc,
		Update: defaultCallBackFunc,
		Delete: defaultCallBackFunc,
		List:   defaultCallBackFunc,
	}

	return nil
}

func NewContainer(c *ContainerConfig) *corev1.Container {
	container := &corev1.Container{
		Name:    c.Name,
		Image:   c.Image,
		Command: c.Command,
		Args:    []string{c.Args},
		Ports:   c.Ports,
		Resources: corev1.ResourceRequirements{
			Limits: corev1.ResourceList{
				corev1.ResourceCPU:              resource.MustParse(getCPUStr(c.Resource.CPU)),
				corev1.ResourceMemory:           resource.MustParse(getByteStr(c.Resource.Mem)),
				corev1.ResourceEphemeralStorage: resource.MustParse(getByteStr(c.Resource.DiskCache)),
			},
			Requests: corev1.ResourceList{
				corev1.ResourceCPU:              resource.MustParse(getCPUStr(c.Resource.CPU / 2.0)),
				corev1.ResourceMemory:           resource.MustParse(getByteStr(c.Resource.Mem / 2)),
				corev1.ResourceEphemeralStorage: resource.MustParse(getByteStr(c.Resource.DiskCache / 2)),
			},
		},
		VolumeMounts:    nil, // TODO: need to fill after Minio Cluster is OK!
		VolumeDevices:   nil, // TODO: need to fill after Minio Cluster is OK!
		ImagePullPolicy: corev1.PullIfNotPresent,
		SecurityContext: c.SecurityContext,
	}

	// GPU Relation
	if c.Resource.BindGPU {
	}
	// Mount Volume(Minio Cluster) Relation
	if c.Resource.BindMount {
		container.Resources.Limits[corev1.ResourceStorage] =
			resource.MustParse(getByteStr(c.Resource.DiskMount))
		container.Resources.Requests[corev1.ResourceStorage] =
			resource.MustParse(getByteStr(c.Resource.DiskMount / 2))
	}

	return container
}

func (c *DeploymentConfig) WithCmdArgs(cmd []string, args string) *DeploymentConfig {
	c.Container.Command = cmd
	c.Container.Args = args
	return c
}

func (c *DeploymentConfig) CancelCmdArgs(ok bool) *DeploymentConfig {
	if ok {
		c.Container.Command = nil
		c.Container.Args = ""
	}
	return c
}

func (c *DeploymentConfig) WithUserHost(uh UserHostConfig) *DeploymentConfig {
	c.Container.UserHost = &uh
	return c
}

func (c *DeploymentConfig) CancelUserHost(ok bool) *DeploymentConfig {
	if ok {
		c.Container.UserHost = nil
	}
	return c
}

func (c *DeploymentConfig) validate() error {
	if c == nil {
		return errors.New("deployment config is null")
	}
	if c.Container == nil {
		return errors.New("container config is null")
	}
	if c.Container.Resource == nil {
		return errors.New("container resource is null")
	}
	return nil
}

func (c *DeploymentConfig) Create() error {
	var createErr error
	c.CallBack.latest, createErr = c.cli.Create(c.ctx, c.app, metav1.CreateOptions{})
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
			c.CallBack.latest, updateErr = c.cli.Update(c.ctx, c.app, metav1.UpdateOptions{})
			return updateErr
		},
	)
	if retryErr != nil {
		return retryErr
	}

	return c.CallBack.Update()
}

func (c *DeploymentConfig) Delete() error {
	deleteErr := c.cli.Delete(c.ctx, c.MetaName, metav1.DeleteOptions{
		PropagationPolicy: util.New(metav1.DeletePropagationForeground),
	})
	if deleteErr != nil {
		return deleteErr
	}

	return c.CallBack.Delete()
}

func (c *DeploymentConfig) List() error {
	var listErr error
	c.CallBack.latest, listErr = c.cli.List(c.ctx, metav1.ListOptions{})
	if listErr != nil {
		return listErr
	}

	return c.CallBack.List()
}

func (c *DeploymentConfig) Stop() (string, error) {
	return c.SaveConfig(), c.Delete()
}

func (c *DeploymentConfig) Restart() error {
	if delErr := c.Delete(); delErr != nil {
		return delErr
	}
	return c.Create()
}

func (c *DeploymentConfig) JSONMarshal() string {
	bs, err := c.app.Marshal()
	if err != nil {
		util.SilentHandleError("marshal error", err)
		return ""
	}
	return util.Bytes2StringNoCopy(bs)
}

// SaveConfig
// after SaveConfig, CallBack is nil
func (c *DeploymentConfig) SaveConfig() string {
	c.CallBack = nil
	bs, err := json.Marshal(c)
	if err != nil {
		util.SilentHandleError("marshal error", err)
		return ""
	}
	return util.Bytes2StringNoCopy(bs)
}

// GetImageURL
// Example:
// 192.168.31.242:8662/kubesphere-io-centos7/haproxy:2.9.6-alpine
func GetImageURL(arti harbor_api.ArtifactURI) string {
	return strings.Join(
		[]string{internalImageRegistry, arti.String()},
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
