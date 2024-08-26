package nginx

import (
	corev1 "k8s.io/api/core/v1"
	"kube/pkg/k8s_api"
	"kube/pkg/util"
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
		ContainerName:      "nginx-web-app",
		ContainerImage:     k8s_api.GetImageURL("nginx:1.14-alpine"),
		ContainerCommand:   []string{"sh", "-c", "while true; do sleep 3600; done"},
		ContainerPort: corev1.ContainerPort{
			Name:          "http",
			HostPort:      8080,
			ContainerPort: 80,
			Protocol:      corev1.ProtocolTCP,
		},
		ContainerResource: &k8s_api.ResourceDecl{
			CPU:       0.5,
			Mem:       256 * util.Mi,
			DiskCache: k8s_api.ContainerLimitDiskCacheDefaultGi * util.Gi,
			//DiskMount: 16 * util.Gi, // none minio cluster bind with s3fs
		},
	}
	err := k8s_api.NewDeployment(&dConf)
	if err != nil {
		panic(err)
	}

	return &dConf
}
