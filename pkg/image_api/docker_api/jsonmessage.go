package docker_api

import (
	"encoding/json"
	"errors"
	"github.com/Juminiy/kube/pkg/log_api/stdlog"
	"github.com/Juminiy/kube/pkg/util"
	"github.com/docker/docker/pkg/jsonmessage"
	"io"
)

//unpack of github.com/docker/docker/pkg/jsonmessage.JSONMessage

// get image ID or full reference from body.body.src.buf JSON
// +example
// {"stream":"Loaded image ID: sha256:d2c94e258dcb3c5ac2798d32e1249e42ef01cba4841c2234249495f87264ac5a\n"}
func getImageIDFromJSONMessage(ioReader io.Reader) string {
	ioJsonDec := json.NewDecoder(ioReader)
	var jsonMessage jsonmessage.JSONMessage
	err := ioJsonDec.Decode(&jsonMessage)
	if err != nil && !errors.Is(err, io.EOF) {
		stdlog.ErrorF("decode json message error: %s", err.Error())
		return ""
	}
	if err0, err1 := jsonMessage.Error != nil, len(jsonMessage.ErrorMessage) != 0; err0 || err1 {
		if err0 {
			stdlog.ErrorF("json message error: %s", jsonMessage.Error.Message)
		}
		if err1 {
			stdlog.ErrorF("json message old error: %s", jsonMessage.ErrorMessage)
		}
		return ""
	}

	messageStreamAttr := jsonMessage.Stream
	if len(messageStreamAttr) == 0 {
		stdlog.Error("json message attribute stream is nil")
		return ""
	}

	//Loaded image ID: sha256:d2c94e258dcb3c5ac2798d32e1249e42ef01cba4841c2234249495f87264ac5a\n
	//Loaded image: hello-world:latest
	return util.StringReplaceAlls(messageStreamAttr, "", "\n", "Loaded image ID: ", "Loaded image: ")
}
