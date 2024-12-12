package docker_api

import (
	"github.com/Juminiy/kube/pkg/util"
	"github.com/docker/docker/api/types/events"
	"github.com/docker/docker/api/types/filters"
	"k8s.io/kube-openapi/pkg/util/sets"
	"time"
)

type ShowEvent struct {
	FromTime time.Time
	ToTime   time.Time
	Type     events.Type
	Filter   Filter
}

func (s ShowEvent) filterArgs() filters.Args {
	filter := filters.NewArgs()
	if len(s.Type) > 0 {
		filter.Add("type", string(s.Type))
	}
	for etyp, eset := range s.Filter {
		for _, ef := range eset.UnsortedList() {
			filter.Add(string(etyp), ef)
		}
	}
	return filter
}

// Filter
// config=<string> config name or ID
// container=<string> container name or ID
// daemon=<string> daemon name or ID
// event=<string> event type
// image=<string> image name or ID
// label=<string> image or container label
// network=<string> network name or ID
// node=<string> node ID
// plugin= plugin name or ID
// scope= local or swarm
// secret=<string> secret name or ID
// service=<string> service name or ID
// type=<string> object to filter by, one of container, image, volume, network, daemon, plugin, node, service, secret or config
// volume=<string> volume name
type Filter map[events.Type]sets.String

const (
	LabelEventType events.Type = "label"
	ScopeEventType events.Type = "scope"
	ScopeLocal                 = "local"
	ScopeSwarm                 = "swarm"
)

func (c *Client) SystemEvent(show ShowEvent) (<-chan events.Message, <-chan error) {
	return c.cli.Events(c.ctx, events.ListOptions{
		Since:   util.ToTimestamp(show.FromTime),
		Until:   util.ToTimestamp(show.ToTime),
		Filters: show.filterArgs(),
	})
}
