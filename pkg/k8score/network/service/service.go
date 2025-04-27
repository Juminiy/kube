package service

import (
	"github.com/Juminiy/kube/pkg/k8score/metadata"
	"github.com/Juminiy/kube/pkg/util"
	corev1 "k8s.io/api/core/v1"
)

func NodePort(
	name, ns string,
	podSelector map[string]string,
) corev1.Service {
	return corev1.Service{
		TypeMeta:   metadata.TypeMetaService(),
		ObjectMeta: metadata.ObjectMeta(name, ns, podSelector, nil),
		Spec: corev1.ServiceSpec{
			Ports:                         []corev1.ServicePort{},
			Selector:                      podSelector,
			ClusterIP:                     "",
			ClusterIPs:                    nil,
			Type:                          corev1.ServiceTypeNodePort,
			ExternalIPs:                   nil,
			SessionAffinity:               "",
			LoadBalancerSourceRanges:      nil,
			ExternalName:                  "",
			ExternalTrafficPolicy:         "",
			HealthCheckNodePort:           0,
			PublishNotReadyAddresses:      false,
			SessionAffinityConfig:         nil,
			IPFamilies:                    nil,
			IPFamilyPolicy:                nil,
			AllocateLoadBalancerNodePorts: nil,
			LoadBalancerClass:             nil,
			InternalTrafficPolicy:         util.New(corev1.ServiceInternalTrafficPolicyLocal),
			TrafficDistribution:           nil,
		},
	}
}
