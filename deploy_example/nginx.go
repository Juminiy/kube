package deploy_example

import (
	corev1 "k8s.io/api/core/v1"
	"kube/kapi"
	"kube/util"
)

// NewNginxDeployment
// Example:
// 1Pod 1Container 1PortExpose
func NewNginxDeployment() *kapi.DeploymentConfig {
	dConf := kapi.DeploymentConfig{
		Namespace:          corev1.NamespaceDefault,
		MetaName:           "nginx-pod-example",
		SpecReplicas:       2,
		SpecSelectorLabels: map[string]string{"app": "nginx-example"},
		SpecTemplateLabels: map[string]string{"app": "nginx-example"},
		ContainerName:      "nginx-web-app",
		ContainerImage:     kapi.GetImageURL("nginx:1.14-alpine"),
		ContainerCommand:   []string{"sh", "-c", "while true; do sleep 3600; done"},
		ContainerPort: corev1.ContainerPort{
			Name:          "http",
			HostPort:      8080,
			ContainerPort: 80,
			Protocol:      corev1.ProtocolTCP,
		},
		ContainerResource: &kapi.ResourceDecl{
			CPU:       0.5,
			Mem:       2 * util.Gi,
			DiskCache: kapi.ContainerLimitDiskCacheDefaultGi * util.Gi,
			//DiskMount: 16 * util.Gi, // none minio cluster bind with s3fs
		},
	}
	err := kapi.NewDeployment(&dConf)
	if err != nil {
		panic(err)
	}

	return &dConf
}
