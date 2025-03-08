package minio_api

import (
	"github.com/Juminiy/kube/pkg/util"
	miniocred "github.com/minio/minio-go/v7/pkg/credentials"
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
	cfgPath, err := os.Open(filepath.Join("testdata", "env", "config.yaml"))
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
		UserID:          321,
		UserName:        "iamhajimao_x",
		BucketQuotaByte: 114514,
		BucketName:      "wwsawqfqf24142121521",
	})
	util.Must(err)
	respJSONStr, err := util.MarshalJSONPretty(resp)
	util.Must(err)
	t.Log(respJSONStr)
}

func TestClient_AtomicDeleteFlow(t *testing.T) {
	err := _cli.AtomicDeleteFlow(Resp{
		Req: Req{
			UserID:          321,
			UserName:        "iamhajimao_x",
			BucketQuotaByte: 114514,
			BucketName:      "wwsawqfqf24142121521",
		},
		CredValue: miniocred.Value{
			AccessKeyID:     "MB4zRO4mP0Rt6ieq9dW0fP688NU5O6sR956s7DN46v9Ar00OM4",
			SecretAccessKey: "oDvb3Auynla8aAqbZjWYV8JX24872JOkx15iAaaFU166oZzVoP5LY6l1H936ae51W6dYCd8EpQEQQEJIBmg64Bgc41l2LQWJwgVX8lgJHATLtXaKmtpfGcG8eSeI06dN",
		},
		CredPolicyName: "321-iamhajimao_x-wwsawqfqf24142121521-k2uO2WhINt7T",
	})
	util.Must(err)
}
