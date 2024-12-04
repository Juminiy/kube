package docker_internal

import (
	"encoding/json"
	"errors"
	"github.com/Juminiy/kube/pkg/log_api/stdlog"
	"github.com/Juminiy/kube/pkg/util"
	"github.com/docker/docker/pkg/jsonmessage"
	"io"
)

//unpack of github.com/docker/docker/pkg/jsonmessage.JSONMessage

// GetImageIDFromImageLoadResp
// get image ID or full reference from body.body.src.buf JSON
// +example
// {"stream":"Loaded image ID: sha256:d2c94e258dcb3c5ac2798d32e1249e42ef01cba4841c2234249495f87264ac5a\n"}
func GetImageIDFromImageLoadResp(ioReader io.Reader) string {
	messageStreamAttr := decodeJSONMessage(ioReader).Stream
	if len(messageStreamAttr) == 0 {
		stdlog.Error("json message attribute stream is nil")
		return ""
	}

	//Loaded image ID: sha256:d2c94e258dcb3c5ac2798d32e1249e42ef01cba4841c2234249495f87264ac5a\n
	//Loaded image: hello-world:latest
	return util.StringReplaceAlls(messageStreamAttr, "", "\n", "Loaded image ID: ", "Loaded image: ")
}

func GetStatusFromImagePushResp(ioReader io.Reader) string {
	messageStatusAttr := decodeJSONMessage(ioReader).Status
	if len(messageStatusAttr) == 0 {
		stdlog.Error("json message attribute status is nil")
		return ""
	}

	return messageStatusAttr
}

func decodeJSONMessage(ioReader io.Reader) (jsonMessage jsonmessage.JSONMessage) {
	ioJsonDec := json.NewDecoder(ioReader)
	err := ioJsonDec.Decode(&jsonMessage)
	if err != nil && !errors.Is(err, io.EOF) {
		stdlog.ErrorF("decode json message error: %s", err.Error())
		return
	}
	//jsonMessageJSONFmt, _ := util.MarshalJSONPretty(jsonMessage)
	//stdlog.ErrorF("raw JSONMessage: %s", jsonMessageJSONFmt)
	if err0, err1 := jsonMessage.Error != nil, len(jsonMessage.ErrorMessage) != 0; err0 || err1 {
		if err0 {
			stdlog.ErrorF("json message error: %s", jsonMessage.Error.Message)
		}
		if err1 {
			stdlog.ErrorF("json message old error: %s", jsonMessage.ErrorMessage)
		}
	}

	return
}
