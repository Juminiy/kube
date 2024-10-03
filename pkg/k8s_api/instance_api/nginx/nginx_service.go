package nginx

import (
	"github.com/Juminiy/kube/pkg/k8s_api/service_api"
	"github.com/Juminiy/kube/pkg/log_api/stdlog"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

func NewService() *service_api.ServiceConfig {
	serviceConfig := &service_api.ServiceConfig{
		Namespace:    corev1.NamespaceDefault,
		MetaName:     "nginx-service-8080-80",
		SpecSelector: map[string]string{"app.kubernetes.io/name": "expose-proxy"},
		SpecPorts: []corev1.ServicePort{
			{
				Name:       "nginx-service-80",
				Protocol:   corev1.ProtocolTCP,
				Port:       8080,
				TargetPort: intstr.FromInt32(80),
			},
		},
	}

	err := service_api.NewService(serviceConfig)
	if err != nil {
		stdlog.ErrorF("nginx service error: %s", err.Error())
	}

	return serviceConfig
}
