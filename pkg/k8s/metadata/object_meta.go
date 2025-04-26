package metadata

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

/*
 * Workloads
 */

func TypeMetaPod() metav1.TypeMeta {
	return metav1.TypeMeta{
		Kind:       "Pod",
		APIVersion: "v1",
	}
}

func TypeMetaDeployment() metav1.TypeMeta {
	return metav1.TypeMeta{
		Kind:       "Deployment",
		APIVersion: "apps/v1",
	}
}

func TypeMetaReplicaSet() metav1.TypeMeta {
	return metav1.TypeMeta{
		Kind:       "ReplicaSet",
		APIVersion: "apps/v1",
	}
}

func TypeMetaStatefulSet() metav1.TypeMeta {
	return metav1.TypeMeta{
		Kind:       "StatefulSet",
		APIVersion: "apps/v1",
	}
}

func TypeMetaDaemonSet() metav1.TypeMeta {
	return metav1.TypeMeta{
		Kind:       "DaemonSet",
		APIVersion: "apps/v1",
	}
}

func TypeMetaJob() metav1.TypeMeta {
	return metav1.TypeMeta{
		Kind:       "Job",
		APIVersion: "batch/v1",
	}
}

func TypeMetaCronJob() metav1.TypeMeta {
	return metav1.TypeMeta{
		Kind:       "CronJob",
		APIVersion: "batch/v1",
	}
}

func TypeMetaReplicationController() metav1.TypeMeta {
	return metav1.TypeMeta{
		Kind:       "ReplicationController",
		APIVersion: "v1",
	}
}

/*
 * Network
 */

func TypeMetaService() metav1.TypeMeta {
	return metav1.TypeMeta{
		Kind:       "v1",
		APIVersion: "Service",
	}
}

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

func TypeMetaEndpointSlice() metav1.TypeMeta {
	return metav1.TypeMeta{
		Kind:       "EndpointSlice",
		APIVersion: "discovery.k8s.io/v1",
	}
}

func TypeMetaNetworkPolicy() metav1.TypeMeta {
	return metav1.TypeMeta{
		Kind:       "NetworkPolicy",
		APIVersion: "networking.k8s.io/v1",
	}
}

/*
 * Storage
 */

func TypeMetaPersistentVolume() metav1.TypeMeta {
	return metav1.TypeMeta{
		Kind:       "PersistentVolume",
		APIVersion: "v1",
	}
}

func TypeMetaPersistentVolumeClaim() metav1.TypeMeta {
	return metav1.TypeMeta{
		Kind:       "PersistentVolumeClaim",
		APIVersion: "v1",
	}
}

func TypeMetaStorageClass() metav1.TypeMeta {
	return metav1.TypeMeta{
		Kind:       "StorageClass",
		APIVersion: "storage.k8s.io/v1",
	}
}

func TypeMetaReferenceGrant() metav1.TypeMeta {
	return metav1.TypeMeta{
		Kind:       "ReferenceGrant",
		APIVersion: "gateway.networking.k8s.io/v1beta1",
	}
}

/*
 * Config
 */

func TypeMetaConfigMap() metav1.TypeMeta {
	return metav1.TypeMeta{
		Kind:       "v1",
		APIVersion: "ConfigMap",
	}
}

func TypeMetaSecret() metav1.TypeMeta {
	return metav1.TypeMeta{
		Kind:       "v1",
		APIVersion: "Secret",
	}
}

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
