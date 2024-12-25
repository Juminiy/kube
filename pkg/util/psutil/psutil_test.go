package psutil

import (
	"github.com/Juminiy/kube/pkg/util"
	"github.com/Juminiy/kube/pkg/util/safe_json"
	"testing"
)

func TestMarshal(t *testing.T) {
	t.Log(util.Bytes2StringNoCopy(MarshalIndent()))
}

func TestGetSysHardV2(t *testing.T) {
	t.Log(safe_json.Pretty(GetSysHardV2()))
}
