package docker_client

func (c *Client) ImagePull(refStr string) (pullResp EventResp, err error) {
	return c.ImageCreate(refStr)
}

type ImagePullResp ImageCreateResp

func (r *EventResp) GetImagePullResp() (resp ImagePullResp) {
	return ImagePullResp(r.GetImageCreateResp())
}
