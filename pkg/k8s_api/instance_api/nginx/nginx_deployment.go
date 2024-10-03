package nginx

import (
	"github.com/Juminiy/kube/pkg/image_api/harbor_api"
	"github.com/Juminiy/kube/pkg/k8s_api/pod_api"
	"github.com/Juminiy/kube/pkg/log_api/stdlog"
	"github.com/Juminiy/kube/pkg/util"
	corev1 "k8s.io/api/core/v1"
)

func NewDeployment() *pod_api.DeploymentConfig {
	dConf := &pod_api.DeploymentConfig{
		Namespace:  corev1.NamespaceDefault,
		MetaName:   "nginx-pod-example",
		MetaLabels: map[string]string{"app.kubernetes.io/name": "expose-proxy"},

		SpecReplicas:       2,
		SpecSelectorLabels: map[string]string{"app": "nginx-example"},
		SpecTemplateLabels: map[string]string{"app": "nginx-example"},

		Container: &pod_api.ContainerConfig{
			Name: "nginx-web-app",
			Image: pod_api.GetImageURL(harbor_api.ArtifactURI{
				Project:    "k8e",
				Repository: "nginx",
				Tag:        "1.14-alpine",
			}),
			Command: []string{"sh", "-c", "while true; do sleep 3600; done"},
			Ports: []corev1.ContainerPort{
				{
					Name:          "http",
					HostPort:      8080,
					ContainerPort: 80,
					Protocol:      corev1.ProtocolTCP},
			},
			Resource: &pod_api.ResourceConfig{
				CPU:       0.5,
				Mem:       256 * util.Mi,
				DiskCache: pod_api.ContainerLimitDiskCacheDefaultGi * util.Gi,
				//DiskMount: 16 * util.Gi, // none minio cluster bind with s3fs
			},
		},
	}

	err := pod_api.NewDeployment(dConf)
	if err != nil {
		stdlog.ErrorF("nginx deployment error: %s", err.Error())
	}

	return dConf
}
