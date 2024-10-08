package harbor_api

import "github.com/Juminiy/kube/pkg/util"

const (
	ProjectCallBack CallBackType = "Project"
	RepoCallBack    CallBackType = "Repository"
)

type (
	CallBackType string
	CallBackOpt  string

	CallBackAttribute struct {
		latest any
		doFunc util.Func
	}
	CallBackAttr  map[CallBackOpt]CallBackAttribute
	CallBackAttrs map[CallBackType]CallBackAttr

	CallBack struct {
		CallBackAttrs
	}
)
