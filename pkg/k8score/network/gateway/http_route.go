package gateway

import (
	"github.com/Juminiy/kube/pkg/k8score/metadata"
	"github.com/Juminiy/kube/pkg/util"
	gatewayv1 "sigs.k8s.io/gateway-api/apis/v1"
)

func NewHTTPRoute() gatewayv1.HTTPRoute {
	return gatewayv1.HTTPRoute{
		TypeMeta:   metadata.TypeMetaHTTPRoute(),
		ObjectMeta: metadata.ObjectMeta("path-route", "", nil, nil),
		Spec: gatewayv1.HTTPRouteSpec{
			CommonRouteSpec: gatewayv1.CommonRouteSpec{},
			Hostnames:       nil,
			Rules: []gatewayv1.HTTPRouteRule{
				{
					Matches: []gatewayv1.HTTPRouteMatch{
						{
							Path: &gatewayv1.HTTPPathMatch{
								Type:  util.New(gatewayv1.PathMatchPathPrefix),
								Value: util.New("/app-kb"),
							},
							Headers: []gatewayv1.HTTPHeaderMatch{
								{
									Type:  util.New(gatewayv1.HeaderMatchExact),
									Name:  "X-GNU-App-Name",
									Value: "app-kb",
								},
								{
									Type:  util.New(gatewayv1.HeaderMatchExact),
									Name:  "X-GNU-App-Tenant-ExternalID",
									Value: "${tenant-external-id}",
								},
							},
						},
					},
					Filters: []gatewayv1.HTTPRouteFilter{},
					BackendRefs: []gatewayv1.HTTPBackendRef{
						{
							BackendRef: gatewayv1.BackendRef{
								BackendObjectReference: gatewayv1.BackendObjectReference{
									Name:      "kb-app-svc-name",
									Namespace: util.New[gatewayv1.Namespace]("tenant-ns"),
									Port:      util.New[gatewayv1.PortNumber](80),
								},
							},
						},
					},
				},
			},
		},
	}
}
