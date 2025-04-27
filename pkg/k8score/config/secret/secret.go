package secret

import (
	"github.com/Juminiy/kube/pkg/k8score/metadata"
	"github.com/Juminiy/kube/pkg/util"
	corev1 "k8s.io/api/core/v1"
)

func ReadOnlyOpaque(
	name, ns string,
	binData map[string][]byte, strData map[string]string,
) corev1.Secret {
	return corev1.Secret{
		TypeMeta:   metadata.TypeMetaSecret(),
		ObjectMeta: metadata.ObjectMeta(name, ns, nil, nil),
		Immutable:  util.New(true),
		Data:       binData,
		StringData: strData,
		Type:       corev1.SecretTypeOpaque,
	}
}
