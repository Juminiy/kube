package volume

import (
	corev1 "k8s.io/api/core/v1"
)

func ConfigMap() corev1.Volume {
	return corev1.Volume{
		Name: "",
		VolumeSource: corev1.VolumeSource{
			ConfigMap: &corev1.ConfigMapVolumeSource{
				LocalObjectReference: corev1.LocalObjectReference{
					Name: "",
				},
				Items:       []corev1.KeyToPath{},
				DefaultMode: nil,
				Optional:    nil,
			},
		},
	}
}
