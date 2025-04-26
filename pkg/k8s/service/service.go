package service

import (
	"github.com/Juminiy/kube/pkg/k8s/metadata"
	"github.com/Juminiy/kube/pkg/util"
	corev1 "k8s.io/api/core/v1"
)

func New() corev1.Service {
	return corev1.Service{
		TypeMeta:   metadata.TypeMetaService(),
		ObjectMeta: metadata.ObjectMeta("", "", nil, nil),
		Spec: corev1.ServiceSpec{
			Ports:                         nil,
			Selector:                      nil,
			ClusterIP:                     "",
			ClusterIPs:                    nil,
			Type:                          "",
			ExternalIPs:                   nil,
			SessionAffinity:               "",
			LoadBalancerIP:                "",
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
