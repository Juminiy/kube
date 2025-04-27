package container

import (
	"github.com/Juminiy/kube/pkg/util"
	"github.com/samber/lo"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

func ProbeHTTPGet(
	schemeHTTPS bool, path, port string, hdr map[string]string,
) *corev1.Probe {
	return &corev1.Probe{
		ProbeHandler: corev1.ProbeHandler{
			HTTPGet: &corev1.HTTPGetAction{
				Path: path,
				Port: intstr.FromString(port),
				Scheme: func() corev1.URIScheme {
					if schemeHTTPS {
						return corev1.URISchemeHTTPS
					}
					return corev1.URISchemeHTTP
				}(),
				HTTPHeaders: lo.MapToSlice(hdr, func(hdrName, hdrValue string) corev1.HTTPHeader {
					return corev1.HTTPHeader{
						Name:  hdrName,
						Value: hdrValue,
					}
				}),
			},
		},
		InitialDelaySeconds:           60,
		TimeoutSeconds:                3 * 60,
		PeriodSeconds:                 10,
		SuccessThreshold:              3,
		FailureThreshold:              3,
		TerminationGracePeriodSeconds: util.New[int64](5 * 60),
	}
}
