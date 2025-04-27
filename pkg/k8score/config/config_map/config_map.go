package config_map

import (
	"github.com/Juminiy/kube/pkg/k8score/metadata"
	"github.com/Juminiy/kube/pkg/util"
	corev1 "k8s.io/api/core/v1"
)

func ReadOnly(
	name, ns string,
	strData map[string]string, binData map[string][]byte,
) corev1.ConfigMap {
	return corev1.ConfigMap{
		TypeMeta:   metadata.TypeMetaConfigMap(),
		ObjectMeta: metadata.ObjectMeta(name, ns, nil, nil),
		Immutable:  util.NewBool(true),
		Data:       strData,
		BinaryData: binData,
	}
}
