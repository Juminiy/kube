package metadata

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func ObjectMeta(
	name, ns string,
	labels, annotations map[string]string,
) metav1.ObjectMeta {
	return metav1.ObjectMeta{
		Name:        name,
		Namespace:   ns,
		Labels:      labels,
		Annotations: annotations,
	}
}
