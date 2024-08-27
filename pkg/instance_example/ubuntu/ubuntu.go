package ubuntu

import (
	corev1 "k8s.io/api/core/v1"
	"kube/pkg/harbor_api"
	"kube/pkg/k8s_api"
	"kube/pkg/util"
	"strings"
)

func NewUbuntuDeployment() *k8s_api.DeploymentConfig {
	dConf := k8s_api.DeploymentConfig{
		Namespace:          corev1.NamespaceDefault,
		MetaName:           "ubuntu22.04-pod-example",
		SpecReplicas:       1,
		SpecSelectorLabels: map[string]string{"app": "example"},
		SpecTemplateLabels: map[string]string{"app": "example"},
		SpecHostNetwork:    false,
		ContainerName:      "ubuntu2204-example",
		ContainerImage: k8s_api.GetImageURL(harbor_api.ArtifactURI{
			Project:    "k8s-images",
			Repository: "ubuntu",
			Tag:        "22.04",
		}),
		ContainerPorts: []corev1.ContainerPort{
			{
				ContainerPort: 22,
				HostPort:      30022,
			},
		},
		ContainerCommand: []string{"/bin/bash", "-c"},
		ContainerArgs: []string{strings.Join(
			[]string{
				"apt-get update && \\",
				"apt-get install -y openssh-server && \\",
				"sed -i 's/PermitRootLogin prohibit-password/PermitRootLogin yes/' /etc/ssh/sshd_config && \\",
				"service ssh start && \\",
				"tail -f /dev/null"},
			"")},
		ContainerResource: &k8s_api.ResourceDecl{
			CPU:       1,
			Mem:       2 * util.Gi,
			DiskCache: k8s_api.ContainerLimitDiskCacheDefaultGi * util.Gi,
			//DiskMount: 16 * util.Gi, // none minio cluster bind with s3fs
		},
		ContainerSecurityContext: &corev1.SecurityContext{
			Capabilities: &corev1.Capabilities{Add: []corev1.Capability{"SYS_TIME"}},
		},
	}
	err := k8s_api.NewDeployment(&dConf)
	if err != nil {
		panic(err)
	}

	return &dConf
}
