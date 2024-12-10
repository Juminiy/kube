package docker_client

import (
	"github.com/distribution/reference"
	"github.com/docker/docker/api/types/registry"
	"net/http"
	"strings"
)

func (c *Client) ImageCreate(parentRefStr string) (createResp EventResp, err error) {
	ref, err := reference.ParseNormalizedNamed(parentRefStr)
	if err != nil {
		return
	}

	return c.imageCreate(map[string]string{
		"fromImage": reference.FamiliarName(ref),
		"tag":       getAPITagFromNamedRef(ref),
	})
}

func (c *Client) imageCreate(queryParam map[string]string) (createResp EventResp, err error) {
	r := c.post("/images/create")
	r.r.SetQueryParams(queryParam).
		SetHeader(registry.AuthHeader, c.reg.Auth)

	resp, err := r.do()
	return *createResp.Parse(resp), err
}

func getAPITagFromNamedRef(ref reference.Named) string {
	if digested, ok := ref.(reference.Digested); ok {
		return digested.Digest().String()
	}
	ref = reference.TagNameOnly(ref)
	if tagged, ok := ref.(reference.Tagged); ok {
		return tagged.Tag()
	}
	return ""
}

/* 1. pull new image
{
    "status": "Pulling from library/hello-world",
    "id": "latest"
}

{
    "status": "Pulling fs layer",
    "progressDetail": {},
    "id": "c1ec31eb5944"
}

{
    "status": "Downloading",
    "progressDetail": {
        "current": 720,
        "total": 2459
    },
    "progress": "[==============>                                    ]     720B/2.459kB",
    "id": "c1ec31eb5944"
}

{
    "status": "Downloading",
    "progressDetail": {
        "current": 2459,
        "total": 2459
    },
    "progress": "[==================================================>]  2.459kB/2.459kB",
    "id": "c1ec31eb5944"
}

{
    "status": "Verifying Checksum",
    "progressDetail": {},
    "id": "c1ec31eb5944"
}

{
    "status": "Download complete",
    "progressDetail": {},
    "id": "c1ec31eb5944"
}

{
    "status": "Extracting",
    "progressDetail": {
        "current": 2459,
        "total": 2459
    },
    "progress": "[==================================================>]  2.459kB/2.459kB",
    "id": "c1ec31eb5944"
}

{
    "status": "Extracting",
    "progressDetail": {
        "current": 2459,
        "total": 2459
    },
    "progress": "[==================================================>]  2.459kB/2.459kB",
    "id": "c1ec31eb5944"
}

{
    "status": "Pull complete",
    "progressDetail": {},
    "id": "c1ec31eb5944"
}

{
    "status": "Digest: sha256:305243c734571da2d100c8c8b3c3167a098cab6049c9a5b066b6021a60fcb966"
}

{
    "status": "Status: Downloaded newer image for hello-world:latest"
}
*/

/* 2. image exists in local
{
    "status": "Pulling from library/hello-world",
    "id": "latest"
}
{
    "status": "Digest: sha256:305243c734571da2d100c8c8b3c3167a098cab6049c9a5b066b6021a60fcb966"
}
{
    "status": "Status: Image is up to date for hello-world:latest"
}
*/

/* 3. image do not exist in registry
{
    "message": "manifest for ???/??? not found: manifest unknown: manifest unknown"
}
*/

type ImageCreateResp struct {
	PullFromRefStr  string
	DownLocalRefStr string
	Digest          string
	NotFound        bool
}

func (r *EventResp) GetImageCreateResp() (resp ImageCreateResp) {
	if r.Status == http.StatusNotFound {
		resp.NotFound = true
		return
	}
	for _, msg := range r.Message {
		if msg == nil {
			return
		}
		statusStr := msg.Status
		if strings.HasPrefix(statusStr, "Pulling from ") {
			resp.PullFromRefStr = strings.TrimPrefix(statusStr, "Pulling from ") + ":" + msg.ID
		}
		if strings.HasPrefix(statusStr, "Digest: ") {
			resp.Digest = strings.TrimPrefix(statusStr, "Digest: ")
		}
		if strings.HasPrefix(statusStr, "Status: Downloaded newer image for ") {
			resp.DownLocalRefStr = strings.TrimPrefix(statusStr, "Status: Downloaded newer image for ")
		} else if strings.HasPrefix(statusStr, "Status: Image is up to date for ") {
			resp.DownLocalRefStr = strings.TrimPrefix(statusStr, "Status: Image is up to date for ")
		}
	}
	return
}
