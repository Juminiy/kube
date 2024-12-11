package docker_client

import "github.com/Juminiy/kube/pkg/log_api/stdlog"

type ImageBuildResp struct {
	Streams  []string `json:"image_build_info"`
	Sha256ID string   `json:"image_sha256_id"`
}

func (r *EventResp) GetImageBuildResp() (resp ImageBuildResp) {
	resp.Streams = make([]string, 0, len(r.Message))
	for _, msg := range r.Message {
		if msg == nil {
			return
		}
		if len(msg.Stream) != 0 {
			resp.Streams = append(resp.Streams, msg.Stream)
		}
		if msg.Aux != nil {
			var auxID = struct {
				ID string
			}{}
			err := DecE(*msg.Aux, &auxID)
			if err != nil {
				stdlog.ErrorF("unmarshal auxID error: %s", err.Error())
			}
		}
	}
	return
}
