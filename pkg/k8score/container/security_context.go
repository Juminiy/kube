package container

import corev1 "k8s.io/api/core/v1"

func SecurityContext() *corev1.SecurityContext {
	return &corev1.SecurityContext{
		Capabilities:             nil,
		Privileged:               nil,
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
	}
}
