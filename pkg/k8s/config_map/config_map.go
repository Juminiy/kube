package config_map

import (
	"github.com/Juminiy/kube/pkg/k8s/metadata"
	"github.com/Juminiy/kube/pkg/util"
	corev1 "k8s.io/api/core/v1"
)

func New(name, ns string) corev1.ConfigMap {
	return corev1.ConfigMap{
		TypeMeta:   metadata.TypeMetaConfigMap(),
		ObjectMeta: metadata.ObjectMeta(name, ns, nil, nil),
		Immutable:  util.NewBool(true),
		Data:       map[string]string{},
		BinaryData: map[string][]byte{},
	}
}
