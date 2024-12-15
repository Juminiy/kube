package docker_api

import (
	"github.com/Juminiy/kube/pkg/util/safe_reflect"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/samber/lo"
	"k8s.io/kube-openapi/pkg/util/sets"
	"k8s.io/utils/set"
)

func (c *Client) ListContainers() ([]types.Container, error) {
	return c.listContainer(_showAll)
}

func (c *Client) ListContainerIds() ([]string, error) {
	crts, err := c.listContainer(_showAll)
	if err != nil {
		return nil, err
	}
	return selectAttr(crts, "ID", func(v any) string {
		return v.(string)
	}), nil
}

func (c *Client) ListContainerNames() ([]string, error) {
	crts, err := c.listContainer(_showAll)
	if err != nil {
		return nil, err
	}
	ss := make([]string, 0, len(crts))
	for i := range crts {
		if len(crts[i].Names) > 0 {
			ss = append(ss, trimPSlash(crts[i].Names[0]))
		}
	}
	return ss, nil
}

// map[any]struct{} -> map[T]struct{} -> []T
func selectAttr[T comparable](crts []types.Container, attrName string, toT func(v any) T) []T {
	return lo.Keys(lo.MapKeys(
		safe_reflect.IndirectOf(crts).SliceStructFieldValues(attrName),
		func(value struct{}, key any) T {
			return toT(key)
		}))
}

// ShowCRT
// show container runtime
type ShowCRT struct {
	// Fuzzy match
	// container from ancestor image
	Image sets.String // ancestor=(<image-name>[:<tag>], <image id>, or <image@digest>)

	// container network
	EPort   sets.String // expose=(<port>[/<proto>]|<startport-endport>/[<proto>])
	PPort   sets.String // publish=(<port>[/<proto>]|<startport-endport>/[<proto>])
	Network sets.String // network=(<network id> or <network name>)

	// container volume
	Volume sets.String // volume=(<volume name> or <mount point destination>)

	// container state
	Health     sets.String  // health=(starting|healthy|unhealthy|none)
	Status     sets.String  // status=(created|restarting|running|removing|paused|exited|dead)
	ExitedCode set.Set[int] // exited=<int> containers with exit code of <int>

	// container start timeline
	Before string // before=(<container id> or <container name>)
	Since  string // since=(<container id> or <container name>)

	Isolation sets.String       // isolation=(default|process|hyperv) (Windows daemon only)
	Task      bool              // is-task=(true|false)
	Label     map[string]string // label=key or label="key=value" of a container label

	// Exact match
	ID   sets.String // id=<ID> a container's ID
	Name sets.String // name=<name> a container's name
}

func (s ShowCRT) filterArgs() filters.Args {
	filter := filters.NewArgs()
	return filter
}

func (c *Client) listContainer(show ShowCRT) ([]types.Container, error) {
	return c.cli.ContainerList(c.ctx,
		container.ListOptions{
			Size:    true,
			All:     true,
			Latest:  false,
			Limit:   c.page.SizeIntValue(),
			Filters: show.filterArgs(),
		})
}

var _showAll = ShowCRT{}
