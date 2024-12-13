package docker_api

import (
	"github.com/Juminiy/kube/pkg/util"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/image"
	"k8s.io/kube-openapi/pkg/util/sets"
	"time"
)

type ClearImage struct {
	Before   time.Time
	LabelEq  sets.String
	LabelNeq sets.String
}

func (c ClearImage) filterArgs() filters.Args {
	filter := filters.NewArgs(
		filters.Arg("dangling", "1"),
		filters.Arg("until", util.ToTimestamp(c.Before)),
	)
	for _, label := range c.LabelEq.UnsortedList() {
		filter.Add("label", label)
	}
	for _, label := range c.LabelNeq.UnsortedList() {
		filter.Add("label!=", label)
	}
	return filter
}

func (c *Client) ImagePrune(clearImage ClearImage) (image.PruneReport, error) {
	return c.cli.ImagesPrune(c.ctx, clearImage.filterArgs())
}

func (c *Client) GC(gc ...util.Func) {}

type HostImageGCFunc util.Func

// HostImageGC
// cli: docker rmi IMAGE_ID
// maybe quota by:
//
// 1. image CREATED: since, before
// 2. image SIZE: bytes(B)
// 3. cache algorithm policy: lru, lfu
// 4. host disk: bytes(B)
func (c *Client) HostImageStorageGC(gcFunc ...HostImageGCFunc) {}
