package pod_api

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/Juminiy/kube/pkg/image_api/harbor_api"
	"github.com/Juminiy/kube/pkg/k8s_api"
	"github.com/Juminiy/kube/pkg/k8s_api/instance_api/cmd_args"
	k8sinternal "github.com/Juminiy/kube/pkg/k8s_api/internal_api"
	"github.com/Juminiy/kube/pkg/log_api/stdlog"
	"github.com/Juminiy/kube/pkg/util"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/client-go/util/retry"
	"strconv"
	"strings"

	_ "k8s.io/client-go/plugin/pkg/client/auth"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	typedappsv1 "k8s.io/client-go/kubernetes/typed/apps/v1"
)

// move it to yaml config
const (
	ContainerRequestResourceRatio    = 0.5 // request = ratio * limit
	ContainerLimitDiskCacheDefaultGi = 8   // ephemeral 8 GiB
)

type (
	DeploymentConfig struct {
		Namespace  string
		MetaName   string
		MetaLabels map[string]string

		SpecReplicas       int32
		SpecSelectorLabels map[string]string
		SpecTemplateLabels map[string]string
		SpecHostNetwork    bool

		Container *ContainerConfig

		cli typedappsv1.DeploymentInterface
		app *appsv1.Deployment
		ctx context.Context
		cbk *k8sinternal.CallBack
	}

	ContainerConfig struct {
		Name            string
		Image           string
		Command         []string
		Args            string
		Ports           []corev1.ContainerPort
		Resource        *ResourceConfig
		SecurityContext *corev1.SecurityContext
		UserHost        *cmd_args.UserHostConfig
	}

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
)

func NewDeployment(c *DeploymentConfig) error {
	if validErr := c.validate(); validErr != nil {
		return validErr
	}

	c.cli = k8s_api.GetClientSet().AppsV1().Deployments(c.Namespace)

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

	c.ctx = util.TODOContext()

	c.cbk = k8sinternal.LogLatestCallBack()

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

func (c *DeploymentConfig) WithUserHost(uh *cmd_args.UserHostConfig) *DeploymentConfig {
	c.Container.UserHost = uh
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
	c.cbk.Latest, c.cbk.LatestErr = c.cli.Create(c.ctx, c.app, metav1.CreateOptions{})
	if c.cbk.LatestErr != nil {
		return c.cbk.LatestErr
	}

	return c.cbk.Create()
}

func (c *DeploymentConfig) Update() error {
	retryErr := retry.RetryOnConflict(
		retry.DefaultRetry,
		func() error {
			c.cbk.Latest, c.cbk.LatestErr = c.cli.Update(c.ctx, c.app, metav1.UpdateOptions{})
			return c.cbk.LatestErr
		},
	)
	if retryErr != nil {
		return retryErr
	}

	return c.cbk.Update()
}

func (c *DeploymentConfig) Delete() error {
	deleteErr := c.cli.Delete(c.ctx, c.MetaName, metav1.DeleteOptions{
		PropagationPolicy: util.New(metav1.DeletePropagationForeground),
	})
	if deleteErr != nil {
		return deleteErr
	}

	return c.cbk.Delete()
}

func (c *DeploymentConfig) List() error {
	c.cbk.Latest, c.cbk.LatestErr = c.cli.List(c.ctx, metav1.ListOptions{})
	if c.cbk.LatestErr != nil {
		return c.cbk.LatestErr
	}

	return c.cbk.List()
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
		stdlog.ErrorF("appsv1 Deployment marshal error: %s", err.Error())
		return ""
	}
	return util.Bytes2StringNoCopy(bs)
}

// SaveConfig
// after SaveConfig, CallBack is nil
func (c *DeploymentConfig) SaveConfig() string {
	c.cbk = nil
	bs, err := json.Marshal(c)
	if err != nil {
		stdlog.ErrorF("deployment config json marshal error: %s", err.Error())
		return ""
	}
	return util.Bytes2StringNoCopy(bs)
}

// GetImageURL
// Example:
// harbor.local:8662/kubesphere-io-centos7/haproxy:2.9.6-alpine
func GetImageURL(arti harbor_api.ArtifactURI) string {
	return strings.Join(
		[]string{k8s_api.GetImageRegistry(), arti.String()},
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
