package metadata

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

/*
 * Networking
 */

func TypeMetaIngress() metav1.TypeMeta {
	return metav1.TypeMeta{
		Kind:       "Ingress",
		APIVersion: "networking.k8s.io/v1",
	}
}

func TypeMetaIngressClass() metav1.TypeMeta {
	return metav1.TypeMeta{
		Kind:       "IngressClass",
		APIVersion: "networking.k8s.io/v1",
	}
}

func TypeMetaNetworkPolicy() metav1.TypeMeta {
	return metav1.TypeMeta{
		Kind:       "NetworkPolicy",
		APIVersion: "networking.k8s.io/v1",
	}
}

/*
 * Gateway
 */

func TypeMetaGateway() metav1.TypeMeta {
	return metav1.TypeMeta{
		Kind:       "Gateway",
		APIVersion: "gateway.networking.k8s.io/v1",
	}
}

func TypeMetaGatewayClass() metav1.TypeMeta {
	return metav1.TypeMeta{
		Kind:       "GatewayClass",
		APIVersion: "gateway.networking.k8s.io/v1",
	}
}

func TypeMetaHTTPRoute() metav1.TypeMeta {
	return metav1.TypeMeta{
		Kind:       "HTTPRoute",
		APIVersion: "gateway.networking.k8s.io/v1",
	}
}

func TypeMetaReferenceGrant() metav1.TypeMeta {
	return metav1.TypeMeta{
		Kind:       "ReferenceGrant",
		APIVersion: "gateway.networking.k8s.io/v1beta1",
	}
}

/*
 * Discovery
 */

func TypeMetaEndpointSlice() metav1.TypeMeta {
	return metav1.TypeMeta{
		Kind:       "EndpointSlice",
		APIVersion: "discovery.k8s.io/v1",
	}
}
