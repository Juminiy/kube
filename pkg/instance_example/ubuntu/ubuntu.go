package ubuntu

import (
	corev1 "k8s.io/api/core/v1"
	"kube/pkg/k8s_api"
	"kube/pkg/util"
)

func NewUbuntuDeployment() *k8s_api.DeploymentConfig {
	dConf := k8s_api.DeploymentConfig{
		Namespace:          corev1.NamespaceDefault,
		MetaName:           "ubuntu22.04-pod-example",
		SpecReplicas:       1,
		SpecSelectorLabels: map[string]string{"app": "example"},
		SpecTemplateLabels: map[string]string{"app": "example"},
		ContainerName:      "ubuntu22.04-example",
		ContainerImage:     k8s_api.GetImageURL("nginx:1.14-alpine"),
		ContainerCommand:   []string{"sh", "-c", "while true; do sleep 3600; done"},
		ContainerResource: &k8s_api.ResourceDecl{
			CPU:       1,
			Mem:       2 * util.Gi,
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
