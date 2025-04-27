package metadata

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// k8s.io/api/core/v1/register.go

/*
 * Workloads
 */

func TypeMetaPod() metav1.TypeMeta {
	return metav1.TypeMeta{
		Kind:       "Pod",
		APIVersion: "v1",
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
		Kind:       "Service",
		APIVersion: "v1",
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

/*
 * Config
 */

func TypeMetaConfigMap() metav1.TypeMeta {
	return metav1.TypeMeta{
		Kind:       "ConfigMap",
		APIVersion: "v1",
	}
}

func TypeMetaSecret() metav1.TypeMeta {
	return metav1.TypeMeta{
		Kind:       "Secret",
		APIVersion: "v1",
	}
}
