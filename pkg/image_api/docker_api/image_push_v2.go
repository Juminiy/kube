package docker_api

import (
	"errors"
	"github.com/Juminiy/kube/pkg/log_api/stdlog"
	"github.com/Juminiy/kube/pkg/util"
	"github.com/Juminiy/kube/pkg/util/safe_json"
	"github.com/distribution/reference"
	"github.com/docker/docker/api/types/registry"
	"github.com/docker/docker/pkg/jsonmessage"
	"github.com/spf13/cast"
	"net/http"
	"strings"
)

var errAPIRespStatus = errors.New("official API response status code is not 200OK")

type PushImageOfficialAPIResp struct {
	Messages []*jsonmessage.JSONMessage
	Aux      []*map[string]any //jsonmessage.JSONMessage .Aux -> []byte -> map[string]any
}

func (r *PushImageOfficialAPIResp) Parse(jsonStrSplitByLineFeed string) {
	jsonStrList := strings.Split(jsonStrSplitByLineFeed, "\r\n")
	r.Messages = make([]*jsonmessage.JSONMessage, len(jsonStrList))
	r.Aux = make([]*map[string]any, len(jsonStrList))
	for i := range jsonStrList {
		msg := jsonmessage.JSONMessage{}
		err := safe_json.GoCCY().Unmarshal(util.String2BytesNoCopy(jsonStrList[i]), &msg)
		if err == nil {
			var msgAux map[string]any
			if msg.Aux != nil {
				err := safe_json.GoCCY().Unmarshal([]byte(*msg.Aux), &msgAux)
				if err == nil {
					r.Aux[i] = &msgAux
				}
			}
			r.Messages[i] = &msg
		} else {
			stdlog.WarnF("jsonmessage[%d], error: %s", i, err.Error())
		}
	}
}

func (r *PushImageOfficialAPIResp) GetDigest() string {
	v := r.getAuxValue("Digest")
	if v != nil {
		return cast.ToString(v)
	}
	return ""
}

func (r *PushImageOfficialAPIResp) GetSize() string {
	v := r.getAuxValue("Size")
	if v != nil {
		return util.MeasureByte(cast.ToInt(v))
	}
	return ""
}

func (r *PushImageOfficialAPIResp) getAuxValue(key string) any {
	for _, auxPtr := range r.Aux {
		if auxPtr != nil {
			if digestValue, ok := util.MapElemOk(*auxPtr, key); ok {
				return digestValue
			}
		}
	}
	return nil
}

func (c *Client) pushImageV2(absRefStr string) (resp PushImageOfficialAPIResp, err error) {
	apiReq := c.restyCli.R().
		SetHeader(registry.AuthHeader, c.reg.Auth)

	refNamed, err := reference.ParseNormalizedNamed(absRefStr)
	if err != nil {
		return resp, err
	}
	pathParamName := reference.FamiliarName(refNamed)
	tagNamed := reference.TagNameOnly(refNamed)
	if tagged, ok := tagNamed.(reference.Tagged); ok {
		apiReq.SetQueryParam("tag", tagged.Tag())
	}
	apiResp, err := apiReq.
		SetPathParam("name", pathParamName).
		Post("/v" + c.version + "/images/{name}/push")
	if err != nil {
		return
	}
	if apiResp.StatusCode() != http.StatusOK {
		return resp, errAPIRespStatus
	}
	resp.Parse(apiResp.String())
	return resp, nil
}
