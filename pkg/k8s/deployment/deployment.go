package deployment

import (
	"github.com/Juminiy/kube/pkg/util"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func New(podName string) appsv1.Deployment {
	return appsv1.Deployment{
		TypeMeta:   metav1.TypeMeta{},
		ObjectMeta: metav1.ObjectMeta{},
		Spec: appsv1.DeploymentSpec{
			Replicas: util.NewInt32(3),
			Selector: nil,
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Name: podName,
				},
				Spec: corev1.PodSpec{
					Volumes:                       nil,
					InitContainers:                nil,
					Containers:                    nil,
					EphemeralContainers:           nil,
					RestartPolicy:                 "",
					TerminationGracePeriodSeconds: nil,
					ActiveDeadlineSeconds:         nil,
					DNSPolicy:                     "",
					NodeSelector:                  nil,
					ServiceAccountName:            "",
					DeprecatedServiceAccount:      "",
					AutomountServiceAccountToken:  nil,
					NodeName:                      "",
					HostNetwork:                   false,
					HostPID:                       false,
					HostIPC:                       false,
					ShareProcessNamespace:         nil,
					SecurityContext:               nil,
					ImagePullSecrets:              nil,
					Hostname:                      "",
					Subdomain:                     "",
					Affinity:                      nil,
					SchedulerName:                 "",
					Tolerations:                   nil,
					HostAliases:                   nil,
					PriorityClassName:             "",
					Priority:                      nil,
					DNSConfig:                     nil,
					ReadinessGates:                nil,
					RuntimeClassName:              nil,
					EnableServiceLinks:            nil,
					PreemptionPolicy:              nil,
					Overhead:                      nil,
					TopologySpreadConstraints:     nil,
					SetHostnameAsFQDN:             nil,
					OS:                            nil,
					HostUsers:                     nil,
					SchedulingGates:               nil,
					ResourceClaims:                nil,
				},
			},
			Strategy:                appsv1.DeploymentStrategy{},
			MinReadySeconds:         0,
			RevisionHistoryLimit:    nil,
			Paused:                  false,
			ProgressDeadlineSeconds: nil,
		},
		Status: appsv1.DeploymentStatus{},
	}
}
