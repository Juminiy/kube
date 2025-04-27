package container

import (
	corev1 "k8s.io/api/core/v1"
)

func New() corev1.Container {
	return corev1.Container{
		Name:                     "",
		Image:                    "",
		Command:                  []string{},
		Args:                     []string{},
		WorkingDir:               "",
		Ports:                    []corev1.ContainerPort{},
		EnvFrom:                  []corev1.EnvFromSource{},
		Env:                      []corev1.EnvVar{},
		Resources:                corev1.ResourceRequirements{},
		ResizePolicy:             []corev1.ContainerResizePolicy{},
		RestartPolicy:            nil,
		VolumeMounts:             []corev1.VolumeMount{},
		VolumeDevices:            []corev1.VolumeDevice{},
		LivenessProbe:            nil,
		ReadinessProbe:           nil,
		StartupProbe:             nil,
		Lifecycle:                nil,
		TerminationMessagePath:   "",
		TerminationMessagePolicy: "",
		ImagePullPolicy:          corev1.PullAlways,
		SecurityContext:          nil,
		Stdin:                    false,
		StdinOnce:                false,
		TTY:                      false,
	}
}
