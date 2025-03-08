package deepseek

import (
	"context"
	"github.com/Juminiy/kube/pkg/util"
	"gopkg.in/yaml.v3"
	"io"
	"os"
	"path/filepath"
	"testing"
)

var _cfg struct {
	DeepSeek struct {
		BaseURL string `yaml:"BaseURL"`
		APIKey  string `yaml:"APIKey"`
	} `yaml:"DeepSeek"`
	OpenAI struct {
		APIKey string `yaml:"APIKey"`
	} `yaml:"OpenAI"`
}
var _cli *Client
var _ctx = context.TODO()

func init() {
	cfgPath, err := os.Open(filepath.Join("token.yaml"))
	util.Must(err)
	cfgBytes, err := io.ReadAll(cfgPath)
	util.Must(err)
	util.Must(yaml.Unmarshal(cfgBytes, &_cfg))
	_cli, err = New(_cfg.DeepSeek.BaseURL, _cfg.DeepSeek.APIKey)
	util.Must(err)
}

func TestCompletions(t *testing.T) {
	resp, err := _cli.Completions("uid:1:some-seek",
		NewCompletionReq("谁是杜甫"))
	util.Must(err)
	t.Log(Pretty(resp))
}

func TestModels(t *testing.T) {
	modelsList, err := _cli.Models()
	util.Must(err)
	t.Log(Pretty(modelsList))
}

func TestBalance(t *testing.T) {
	myBalance, err := _cli.Balance()
	util.Must(err)
	t.Log(Pretty(myBalance))
}
