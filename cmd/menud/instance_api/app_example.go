package instance_api

import (
	"github.com/Juminiy/kube/pkg/k8s_api/instance_api/nginx"
	"github.com/Juminiy/kube/pkg/util"
)

func AppMenu(app, action string) {

	var (
		nginxApp *nginx.App
	)

	switch app {

	case "nginx":
		nginxApp = nginx.DefaultApp()

		switch action {
		case "create":
			util.SilentError(nginxApp.Deployment.Create())
			util.SilentError(nginxApp.Service.Create())

		case "delete":
			util.SilentError(nginxApp.Deployment.Delete())
			util.SilentError(nginxApp.Service.Delete())

		case "update":
			util.SilentError(nginxApp.Deployment.Update())
			util.SilentError(nginxApp.Service.Update())

		case "list":
			util.SilentError(nginxApp.Deployment.List())
			util.SilentError(nginxApp.Service.List())

		}

	}

}
