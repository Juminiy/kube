package volume

import corev1 "k8s.io/api/core/v1"

func ROConfigMap(configMapName string) corev1.Volume {
	return corev1.Volume{
		Name: configMapName,
		VolumeSource: corev1.VolumeSource{
			ConfigMap: &corev1.ConfigMapVolumeSource{
				LocalObjectReference: corev1.LocalObjectReference{},
				Items:                nil,
				DefaultMode:          nil,
				Optional:             nil,
			},
		},
	}
}
