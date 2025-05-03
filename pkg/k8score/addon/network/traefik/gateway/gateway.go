package gateway

import (
	"github.com/Juminiy/kube/pkg/k8score/metadata"
	gatewayv1 "sigs.k8s.io/gateway-api/apis/v1"
)

func ClusterClass() gatewayv1.GatewayClass {
	return gatewayv1.GatewayClass{
		TypeMeta:   metadata.TypeMetaGatewayClass(),
		ObjectMeta: metadata.ObjectMeta("traefik-gateway-class", "", nil, nil),
		Spec: gatewayv1.GatewayClassSpec{
			ControllerName: "traefik.io/gateway-controller",
		},
	}
}

func NewGateway() gatewayv1.Gateway {
	return gatewayv1.Gateway{
		TypeMeta:   metadata.TypeMetaGateway(),
		ObjectMeta: metadata.ObjectMeta("", "", nil, nil),
		Spec: gatewayv1.GatewaySpec{
			GatewayClassName: "traefik-gateway-class",
		},
	}
}
