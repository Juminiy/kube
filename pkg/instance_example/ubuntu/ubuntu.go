package ubuntu

import (
	corev1 "k8s.io/api/core/v1"
	"kube/pkg/harbor_api"
	"kube/pkg/instance_example/cmd_args"
	"kube/pkg/k8s_api"
	"kube/pkg/util"
)

func NewUbuntuDeployment() *k8s_api.DeploymentConfig {
	dConf := &k8s_api.DeploymentConfig{
		Namespace:          corev1.NamespaceDefault,
		MetaName:           "ubuntu-pro-2204",
		SpecReplicas:       1,
		SpecSelectorLabels: map[string]string{"app": "example"},
		SpecTemplateLabels: map[string]string{"app": "example"},
		SpecHostNetwork:    false,
		ContainerName:      "ubuntu-pro-2204",
		ContainerImage: k8s_api.GetImageURL(harbor_api.ArtifactURI{
			Project:    "library",
			Repository: "ubuntu-s",
			Tag:        "22.04",
		}),
		ContainerPorts: []corev1.ContainerPort{
			{
				ContainerPort: 22,
				HostPort:      30022,
			},
		},
		UserServerDecl: k8s_api.UserServerDecl{
			HostName: "hln",
			UserName: "hln",
			Password: "hln@666",
		},
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
	// exec cmd args /bin/bash
	dConf.WithCmdArgs(
		[]string{"/bin/bash", "-c"},
		//args[0]
		cmd_args.ArgsConcat(
			cmd_args.UbuntuUpdateUpgrade,
			cmd_args.UbuntuInstall("systemd"),
			//cmd_args.LinuxHostNameCtl(dConf.HostName),
			//cmd_args.LinuxAddUser(dConf.UserName),
			//cmd_args.LinuxSetUserPassword(dConf.Password),
			cmd_args.UbuntuInstall("openssh-server"),
			cmd_args.LinuxPermitSSHLoginByRoot,
			cmd_args.LinuxServiceStart("ssh"),
			cmd_args.LinuxTerminalAlwaysOpen),
	).CancelCmd(false)
	err := k8s_api.NewDeployment(dConf)
	if err != nil {
		panic(err)
	}

	return dConf
}
