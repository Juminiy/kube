package ubuntu

import (
	"github.com/Juminiy/kube/pkg/image_api/harbor_api"
	"github.com/Juminiy/kube/pkg/k8s_api/instance_api/internal/cmd_args"
	"github.com/Juminiy/kube/pkg/k8s_api/pod_api"
	"github.com/Juminiy/kube/pkg/log_api/stdlog"
	"github.com/Juminiy/kube/pkg/util"
	corev1 "k8s.io/api/core/v1"
)

func NewUbuntuDeployment() *pod_api.DeploymentConfig {
	dConf := &pod_api.DeploymentConfig{
		Namespace:          corev1.NamespaceDefault,
		MetaName:           "ubuntu-pro-2204",
		SpecReplicas:       1,
		SpecSelectorLabels: map[string]string{"app": "example"},
		SpecTemplateLabels: map[string]string{"app": "example"},
		SpecHostNetwork:    false,
		Container: &pod_api.ContainerConfig{
			Name: "ubuntu-pro-2204",
			Image: pod_api.GetImageURL(harbor_api.ArtifactURI{
				Project:    "library",
				Repository: "ubuntu-s",
				Tag:        "22.04",
			}),
			Ports: []corev1.ContainerPort{
				{
					ContainerPort: 22,
					HostPort:      30022,
				},
			},
			Resource: &pod_api.ResourceConfig{
				CPU:       1,
				Mem:       2 * util.Gi,
				DiskCache: pod_api.ContainerLimitDiskCacheDefaultGi * util.Gi,
				//DiskMount: 16 * util.Gi, // none minio cluster bind with s3fs
			},
			SecurityContext: &corev1.SecurityContext{
				Capabilities: &corev1.Capabilities{Add: []corev1.Capability{"SYS_TIME"}},
				Privileged:   util.NewBool(true),
			},
		},
	}

	// exec cmd args /bin/bash
	dConf.
		WithCmdArgs(
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
			//cmd_args.S3fsMount{
			//	Key:        "",
			//	Dir:        "",
			//	BucketName: "",
			//	MinioAddr:  "",
			//}.Args(),
		).CancelCmdArgs(false).
		WithUserHost(pod_api.UserHostConfig{
			HostName: "hln",
			UserName: "hln",
			Password: "hln@666",
		}).CancelUserHost(true)

	err := pod_api.NewDeployment(dConf)
	if err != nil {
		stdlog.ErrorF("ubuntu deployment error: %s", err.Error())
	}

	return dConf
}
