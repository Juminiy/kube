package nginx

import (
	"github.com/Juminiy/kube/pkg/k8s_api/pod_api"
	"github.com/Juminiy/kube/pkg/k8s_api/service_api"
)

type App struct {
	Deployment *pod_api.DeploymentConfig
	Service    *service_api.ServiceConfig
}

func DefaultApp() *App {
	return &App{
		Deployment: NewDeployment(),
		Service:    NewService(),
	}
}
