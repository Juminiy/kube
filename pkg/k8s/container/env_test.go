package container

import (
	"github.com/Juminiy/kube/pkg/util"
	"gopkg.in/yaml.v3"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	"testing"
)

func TestContainerEnv(t *testing.T) {
	bs, err := yaml.Marshal(corev1.Container{
		SecurityContext: &corev1.SecurityContext{
			Capabilities:             nil,
			Privileged:               util.New(true),
			SELinuxOptions:           nil,
			WindowsOptions:           nil,
			RunAsUser:                nil,
			RunAsGroup:               nil,
			RunAsNonRoot:             nil,
			ReadOnlyRootFilesystem:   nil,
			AllowPrivilegeEscalation: nil,
			ProcMount:                nil,
			SeccompProfile:           nil,
			AppArmorProfile:          nil,
		},
		Resources: corev1.ResourceRequirements{
			Limits: corev1.ResourceList{
				corev1.ResourceCPU:    resource.MustParse("1000m"),
				corev1.ResourceMemory: resource.MustParse("256Mi"),
			},
			Requests: corev1.ResourceList{
				corev1.ResourceCPU:    resource.MustParse("500m"),
				corev1.ResourceMemory: resource.MustParse("128Mi"),
			},
		},
		Env: []corev1.EnvVar{
			{
				Name:  "MINIO_ROOT_USER",
				Value: "AKIAIOSFODNN7EXAMPLE",
			},
			{
				Name:  "MINIO_ROOT_PASSWORD",
				Value: "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY",
			},
		},
		EnvFrom: nil,
	})
	if err != nil {
		t.Error(err)
	}
	t.Logf("\n%s", bs)
}
