package docker_api

import (
	"github.com/Juminiy/kube/pkg/util"
	"github.com/docker/docker/api/types/registry"
	"gopkg.in/yaml.v3"
	"io"
	"os"
	"path/filepath"
)

var _cli *Client
var _cfg struct {
	Docker struct {
		Addr    string `yaml:"addr"`
		Version string `yaml:"version"`
	} `yaml:"docker"`
	Registry struct {
		Addr     string `yaml:"addr"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	} `yaml:"registry"`
}

func init() {
	cfgPath, err := os.Open(filepath.Join("testdata", "env", "bupt_vpn.yaml"))
	util.Must(err)
	cfgBytes, err := io.ReadAll(cfgPath)
	util.Must(err)
	util.Must(yaml.Unmarshal(cfgBytes, &_cfg))
	_cli, err = New(_cfg.Docker.Addr, _cfg.Docker.Version)
	util.Must(err)
	_cli.WithRegistryAuth(&registry.AuthConfig{
		Username:      _cfg.Registry.Username,
		Password:      _cfg.Registry.Password,
		ServerAddress: _cfg.Registry.Addr,
	}).
		WithProject("library")
}

var _testTar = filepath.Join("testdata", "tar_gz")
