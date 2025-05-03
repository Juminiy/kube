package gateway

import (
	"github.com/Juminiy/kube/pkg/k8score/metadata"
	"github.com/Juminiy/kube/pkg/util"
	gatewayv1 "sigs.k8s.io/gateway-api/apis/v1"
)

func NewGateway() gatewayv1.Gateway {
	return gatewayv1.Gateway{
		TypeMeta:   metadata.TypeMetaGateway(),
		ObjectMeta: metadata.ObjectMeta("", "", nil, nil),
		Spec: gatewayv1.GatewaySpec{
			GatewayClassName: "app-gw",
			Listeners: []gatewayv1.Listener{
				{
					Name:     "app-gw-listener",
					Port:     gatewayv1.PortNumber(80),
					Protocol: gatewayv1.HTTPProtocolType,
					AllowedRoutes: &gatewayv1.AllowedRoutes{
						Namespaces: &gatewayv1.RouteNamespaces{
							From: util.New(gatewayv1.NamespacesFromSame),
						},
						Kinds: []gatewayv1.RouteGroupKind{},
					},
				},
			},
			Addresses:        nil,
			Infrastructure:   nil,
			BackendTLS:       nil,
			AllowedListeners: nil,
		},
	}
}
