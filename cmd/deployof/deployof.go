package main

import (
	"flag"
	"fmt"
	"github.com/Juminiy/kube/cmd/deployof/handlers"
	"github.com/Juminiy/kube/cmd/deployof/sdk"
	"github.com/Juminiy/kube/pkg/image_api/docker_api/docker_inst"
	"github.com/Juminiy/kube/pkg/storage_api/minio_api/minio_inst"
	"github.com/Juminiy/kube/pkg/util"
	"github.com/docker/docker/api/types/registry"
	"github.com/gin-gonic/gin"
	fiberlog "github.com/gofiber/fiber/v3/log"
	"github.com/minio/madmin-go/v3"
	"github.com/minio/minio-go/v7"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

func main() {}

func init() {
	flag.StringVar(&_configPath, "config", "", "config to start app")
	flag.Parse()
	util.SeqRun(InitConfig, InitMinio, InitDocker)

	app := gin.Default()
	app.Use(gin.Recovery())
	apiv1 := app.Group("/api/v1")
	apiv1.POST("/image/upload", handlers.ImageUpload)
	apiv1.GET("/image/download", handlers.ImageDownloadURL)
	log.Fatal(app.Run(fmt.Sprintf("%s:%d", _config.Web.Host, _config.Web.Port)))
}

var (
	_configPath string
	_config     struct {
		Web struct {
			Host string `yaml:"host"`
			Port int    `yaml:"port"`
		}
		Docker struct {
			Addr     string `yaml:"addr"`
			Version  string `yaml:"version"`
			Registry struct {
				Addr     string `yaml:"addr"`
				Username string `yaml:"username"`
				Password string `yaml:"password"`
			}
		}
		Minio struct {
			Addr            string `yaml:"addr"`
			AccessKeyID     string `yaml:"accessKeyID"`
			SecretAccessKey string `yaml:"secretAccessKey"`
			Public          string `yaml:"public"`
		}
	}
)

func InitConfig() {
	fileBytes, err := os.ReadFile(_configPath)
	if err != nil {
		log.Fatal(fmt.Errorf("read yaml file path: %s error: %s", _configPath, err))
	}

	err = yaml.Unmarshal(fileBytes, &_config)
	if err != nil {
		log.Fatal(fmt.Errorf("unmarshal yaml file bytes: %v, yaml config instance: %#v, error: %s", fileBytes, _config, err.Error()))
	}
}

func InitDocker() {
	docker_inst.New().
		WithHost(_config.Docker.Addr).
		WithVersion(_config.Docker.Version).
		Load()
	docker_inst.WithRegistryAuth(&registry.AuthConfig{
		Username:      _config.Docker.Registry.Username,
		Password:      _config.Docker.Registry.Password,
		ServerAddress: _config.Docker.Registry.Addr,
	})
	sdk.Init(_config.Docker.Registry.Username)
	fiberlog.Info("docker init success")
}

func InitMinio() {
	minio_inst.New().
		WithEndpoint(_config.Minio.Addr).
		WithAccessKeyID(_config.Minio.AccessKeyID).
		WithSecretAccessKey(_config.Minio.SecretAccessKey).
		WithPublicBucket(_config.Minio.Public).
		WithOpts(minio.Options{}, madmin.Options{}).
		LoadOpts()
	fiberlog.Info("minio init success")
}
