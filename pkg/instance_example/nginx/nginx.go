package nginx

import (
	"github.com/Juminiy/kube/pkg/harbor_api"
	"github.com/Juminiy/kube/pkg/k8s_api"
	"github.com/Juminiy/kube/pkg/util"
	corev1 "k8s.io/api/core/v1"
)

// NewNginxDeployment
// Example:
// 1Pod 1Container 1PortExpose
func NewNginxDeployment() *k8s_api.DeploymentConfig {
	dConf := k8s_api.DeploymentConfig{
		Namespace:          corev1.NamespaceDefault,
		MetaName:           "nginx-pod-example",
		SpecReplicas:       2,
		SpecSelectorLabels: map[string]string{"app": "nginx-example"},
		SpecTemplateLabels: map[string]string{"app": "nginx-example"},
		Container: &k8s_api.ContainerConfig{
			Name: "nginx-web-app",
			Image: k8s_api.GetImageURL(harbor_api.ArtifactURI{
				Project:    "kubesphere-io-centos7",
				Repository: "nginx",
				Tag:        "1.14-alpine",
			}),
			Command: []string{"sh", "-c", "while true; do sleep 3600; done"},
			Ports: []corev1.ContainerPort{
				corev1.ContainerPort{
					Name:          "http",
					HostPort:      8080,
					ContainerPort: 80,
					Protocol:      corev1.ProtocolTCP},
			},
			Resource: &k8s_api.ResourceConfig{
				CPU:       0.5,
				Mem:       256 * util.Mi,
				DiskCache: k8s_api.ContainerLimitDiskCacheDefaultGi * util.Gi,
				//DiskMount: 16 * util.Gi, // none minio cluster bind with s3fs
			},
		},
	}
	err := k8s_api.NewDeployment(&dConf)
	if err != nil {
		panic(err)
	}

	return &dConf
}
