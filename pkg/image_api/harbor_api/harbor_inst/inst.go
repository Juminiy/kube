package harbor_inst

import (
	"github.com/Juminiy/kube/pkg/image_api/harbor_api"
	"github.com/Juminiy/kube/pkg/log_api/stdlog"
)

// global config
var (
	_harborRegistry string
	_harborInsecure bool
	_harborUsername string
	_harborPassword string
)

// global var
var (
	_harborClient *harbor_api.Client
)

func Init() {
	var harborClientError error
	_harborClient, harborClientError = harbor_api.New(
		_harborRegistry,
		_harborInsecure,
		_harborUsername,
		_harborPassword,
	)
	if harborClientError != nil {
		stdlog.ErrorF("harbor registry: %s, insecure: %v, username: %s, password: ******, error: %s",
			_harborRegistry, _harborInsecure, _harborUsername, harborClientError.Error())
		return
	}

}
