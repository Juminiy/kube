package metadata

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// k8s.io/api/batch/v1/register.go

/*
 * Workloads
 */

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
