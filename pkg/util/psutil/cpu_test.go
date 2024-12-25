package psutil

import (
	"github.com/Juminiy/kube/pkg/util/safe_json"
	"testing"
)

func TestCPUV2(t *testing.T) {
	t.Log(safe_json.Pretty(cpuV2()))
}
