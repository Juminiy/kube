package harbor_api

import (
	"github.com/Juminiy/kube/pkg/util"
	"gopkg.in/yaml.v3"
	"io"
	"os"
	"path/filepath"
)

var _cli *Client
var _cfg struct {
	Addr     string `yaml:"addr"`
	Insecure bool   `yaml:"insecure"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

func init() {
	cfgPath, err := os.Open(filepath.Join("testdata", "env", "taveen_wz.yaml"))
	util.Must(err)
	cfgBytes, err := io.ReadAll(cfgPath)
	util.Must(err)
	util.Must(yaml.Unmarshal(cfgBytes, &_cfg))
	_cli, err = New(_cfg.Addr, _cfg.Insecure, _cfg.Username, _cfg.Password)
	util.Must(err)
}
