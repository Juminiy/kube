package minio_api

import (
	"github.com/Juminiy/kube/pkg/log_api/stdlog"
	"github.com/Juminiy/kube/pkg/util"
	"gopkg.in/yaml.v3"
	"io"
	"os"
	"path/filepath"
	"testing"
)

var _cli *Client
var _cfg struct {
	Endpoint        string `yaml:"endpoint"`
	AccessKeyID     string `yaml:"accessKeyID"`
	SecretAccessKey string `yaml:"secretAccessKey"`
	Secure          bool   `yaml:"secure"`
}

func init() {
	cfgPath, err := os.Open(filepath.Join("testdata", "env", "taveen_rdev.yaml"))
	util.Must(err)
	cfgBytes, err := io.ReadAll(cfgPath)
	util.Must(err)
	util.Must(yaml.Unmarshal(cfgBytes, &_cfg))
	_cli, err = New(_cfg.Endpoint, _cfg.AccessKeyID, _cfg.SecretAccessKey, "", _cfg.Secure)
	util.Must(err)
}

// +passed
func TestClient_AtomicWorkflow(t *testing.T) {
	resp, err := _cli.AtomicWorkflow(Req{
		UserID:          16,
		UserName:        "calimimicc",
		BucketQuotaByte: 114514,
		BucketName:      "funkqqqzaaweqaadfafwq4124asf1122",
	})
	if err != nil {
		panic(err)
	}
	respJSONStr, err := util.MarshalJSONPretty(resp)
	if err != nil {
		panic(err)
	}
	stdlog.Info(respJSONStr)
}
