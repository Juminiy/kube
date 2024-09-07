package main

import (
	"flag"
	"fmt"
	"github.com/Juminiy/kube/cmd/menud/config"
	harbormenu "github.com/Juminiy/kube/cmd/menud/harbor_api"
	instancemenu "github.com/Juminiy/kube/cmd/menud/instance_api"
	k8smenu "github.com/Juminiy/kube/cmd/menud/k8s_api"
	"github.com/Juminiy/kube/pkg/image_api/docker_api/docker_inst"
	"github.com/Juminiy/kube/pkg/image_api/harbor_api/harbor_inst"
	"github.com/Juminiy/kube/pkg/k8s_api"
	"github.com/Juminiy/kube/pkg/log_api/stdlog"
	"github.com/Juminiy/kube/pkg/log_api/zaplog"
	"github.com/Juminiy/kube/pkg/storage_api/minio_api/minio_inst"
	"github.com/Juminiy/kube/pkg/util"
	ldversion "github.com/Juminiy/kube/version"
	"gopkg.in/yaml.v3"
	"os"
)

func main() {
	initFlag()
	ldversion.Info(version)
	initGlobalConfig()

	util.SeqRun(
		initLog,
		initHarbor,
		initDocker,
		initKubernetes,
		initMinio,
	)

	var (
		setting                   string
		moduleOf, appOf, actionOf string
	)

	for {
		fmt.Printf("setting [help | quit | next]: ")
		if _, err := fmt.Scanf("%s", &setting); err != nil {
			fmt.Printf("error setting: %v\n", err)
		}
		if helpMenu(setting) == helpRetCode {
			continue
		}
		fmt.Printf("module [cluster | deploy | harbor]: ")
		if _, err := fmt.Scanf("%s %s %s", &moduleOf, &appOf, &actionOf); err != nil {
			fmt.Printf("error input: %v\n", err)
			return
		}
		switch moduleOf {
		case "cluster":
			k8smenu.Menu(appOf, actionOf)
		case "deploy":
			instancemenu.Menu(appOf, actionOf)
		case "harbor":
			harbormenu.Menu(appOf, actionOf)
		}
	}

}

const (
	helpRetCode int8 = 0
	nextRetCode int8 = 1
)

func helpMenu(s ...string) int8 {
	if s[0] == "help" || s[0] == "h" {
		fmt.Println("help | quit | none [module | app | action]")
		fmt.Println("help: help | h")
		fmt.Println("quit: quit | q")
		fmt.Println("next: next | n")
		fmt.Println("module: cluster | deploy | harbor")
		fmt.Println("app: [log | node | deployment] | [nginx | ubuntu]")
		fmt.Println("action: [create | update | delete | list | start-sync | stop-sync]")
		return helpRetCode
	} else if s[0] == "quit" || s[0] == "q" {
		os.Exit(0)
	}
	return nextRetCode
}

// global Flags
var (
	version        *bool
	configYamlPath *string
)

func initFlag() {
	version = flag.Bool("v", false, "print version json info")
	configYamlPath = flag.String("c", "", "menud config file")
	flag.Parse()
}

var (
	_globalConfig config.Config
)

func initGlobalConfig() {
	fileBytes, err := os.ReadFile(*configYamlPath)
	if err != nil {
		stdlog.ErrorF("read yaml file path: %s error: %s", *configYamlPath, err)
	}

	err = yaml.Unmarshal(fileBytes, &_globalConfig)
	if err != nil {
		stdlog.ErrorF("unmarshal yaml file bytes: %v, yaml config instance: %#v, error: %s", fileBytes, _globalConfig, err.Error())
	}
	stdlog.Info("global config init success")
}

func initLog() {
	zaplog.NewConfig().
		WithLogEngine(_globalConfig.Log.Engine).
		WithLogLevel(_globalConfig.Log.Zap.Level).
		WithLogCaller(_globalConfig.Log.Zap.Caller).
		WithLogStackTrace(_globalConfig.Log.Zap.Stacktrace).
		WithOutputPaths(_globalConfig.Log.Zap.Path...).
		WithErrorOutputPaths(_globalConfig.Log.Zap.InternalPath...).
		Load()
	stdlog.Info("zaplog init success")
}

func initHarbor() {
	harbor_inst.NewConfig().
		WithRegistry(_globalConfig.Harbor.Registry).
		WithUsername(_globalConfig.Harbor.Username).
		WithPassword(_globalConfig.Harbor.Password).
		Load()
	stdlog.Info("harbor client init success")
}

func initDocker() {
	docker_inst.NewConfig().
		WithHost(_globalConfig.Docker.Host).
		WithVersion(_globalConfig.Docker.Version).
		Load()
	stdlog.Info("docker client init success")
}

func initKubernetes() {
	k8s_api.WithKubeConfigPath(_globalConfig.Kubernetes.KubeConfigPath)
	k8s_api.WithImageRegistry(_globalConfig.Harbor.Registry)
	k8s_api.Load()
	stdlog.Info("kubernetes client init success")
}

func initMinio() {
	minio_inst.NewConfig().
		WithEndpoint(_globalConfig.Minio.Endpoint).
		WithAccessKeyID(_globalConfig.Minio.AccessKeyID).
		WithSecretAccessKey(_globalConfig.Minio.SecretAccessKey).
		Load()
	stdlog.Info("minio client init success")
}
