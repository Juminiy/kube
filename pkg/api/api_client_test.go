package api

import (
	"github.com/Juminiy/kube/pkg/util"
	"gopkg.in/yaml.v3"
	"io"
	"os"
	"path/filepath"
)

var _cli *Client
var _cfg struct {
	AttrA string
	AttrB string
	AttrC bool
}

func init() {
	cfgPath, err := os.Open(filepath.Join("testdata", "env", "env.yaml"))
	util.Must(err)
	cfgBytes, err := io.ReadAll(cfgPath)
	util.Must(err)
	util.Must(yaml.Unmarshal(cfgBytes, &_cfg))
	_cli, err = New(_cfg.AttrA, _cfg.AttrB, _cfg.AttrC)
	util.Must(err)
}
