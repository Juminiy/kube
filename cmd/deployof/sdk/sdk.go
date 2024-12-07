package sdk

import (
	"github.com/Juminiy/kube/pkg/image_api/docker_api/docker_inst"
	"github.com/Juminiy/kube/pkg/storage_api/minio_api/minio_inst"
	"github.com/docker/docker/api/types/registry"
	"github.com/gofiber/fiber/v3/log"
)

func init() {
	minio_inst.New().
		WithEndpoint("192.168.3.37:9000").
		WithAccessKeyID("chisato").
		WithSecretAccessKey("lyy001202").
		WithPublicBucket("library").
		Load()
	log.Info("minio init success")

	docker_inst.New().
		WithHost("tcp://127.0.0.1:2375").
		WithVersion("1.47").
		Load()
	docker_inst.WithRegistryAuth(&registry.AuthConfig{
		Username:      "hln7897",
		Password:      "LinYuanYuanSuKiDa001202",
		ServerAddress: "", // docker registry
	})
	log.Info("docker init success")
}
