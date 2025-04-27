package metadata

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// k8s.io/api/apps/v1/register.go

/*
 * Workloads
 */

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

func TypeMetaControllerRevision() metav1.TypeMeta {
	return metav1.TypeMeta{
		Kind:       "ControllerRevision",
		APIVersion: "apps/v1",
	}
}
