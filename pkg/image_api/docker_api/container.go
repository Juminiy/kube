package docker_api

import (
	"github.com/Juminiy/kube/pkg/util"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	k8sutilsets "k8s.io/apimachinery/pkg/util/sets"
	"strings"
)

func (c *Client) ListContainerFullIds() ([]string, error) {
	resultMap, err := c.listContainerWithFields("Id")
	return resultMap["Id"], err
}

func (c *Client) ListContainerIds() ([]string, error) {
	resultMap, err := c.listContainerWithFields("Id")

	ids := make([]string, len(resultMap["Id"]))
	for i, id := range resultMap["Id"] {
		ids[i] = id[:12]
	}

	return ids, err
}

func (c *Client) ListContainerNames() ([]string, error) {
	resultMap, err := c.listContainerWithFields("Name")
	return resultMap["Name"], err
}

func (c *Client) listContainerWithFields(fields ...string) (map[string][]string, error) {
	fieldSet := k8sutilsets.New[string](fields...)
	containers, err := c.ListContainers()
	if err != nil {
		return nil, err
	}
	resultSet := make(map[string][]string, len(fields))
	if fieldSet.Has("Id") {
		for _, con := range containers {
			resultSet["Id"] = append(resultSet["Id"], con.ID)
		}
	}

	if fieldSet.Has("Name") {
		for _, con := range containers {
			resultSet["Name"] = append(resultSet["Name"], removePrefixSlash(con.Names[0]))
		}
	}

	if fieldSet.Has("Image") {
		for _, con := range containers {
			resultSet["Image"] = append(resultSet["Image"], con.Image)
		}
	}

	if fieldSet.Has("ImageID") {
		for _, con := range containers {
			resultSet["ImageID"] = append(resultSet["ImageID"], con.ImageID)
		}
	}

	if fieldSet.Has("Command") {
		for _, con := range containers {
			resultSet["Command"] = append(resultSet["Command"], con.Command)
		}
	}

	if fieldSet.Has("Created") {
		for _, con := range containers {
			resultSet["Created"] = append(resultSet["Created"], util.I64toa(con.Created))
		}
	}

	if fieldSet.Has("SizeRw") {
		for _, con := range containers {
			resultSet["SizeRw"] = append(resultSet["SizeRw"], util.I64toa(con.SizeRw))
		}
	}

	if fieldSet.Has("SizeRootFs") {
		for _, con := range containers {
			resultSet["SizeRootFs"] = append(resultSet["SizeRootFs"], util.I64toa(con.SizeRootFs))
		}
	}

	if fieldSet.Has("State") {
		for _, con := range containers {
			resultSet["State"] = append(resultSet["State"], con.State)
		}
	}

	if fieldSet.Has("Status") {
		for _, con := range containers {
			resultSet["Status"] = append(resultSet["Status"], con.Status)
		}
	}

	return resultSet, nil
}

func (c *Client) ListContainers() ([]types.Container, error) {
	return c.cli.ContainerList(
		c.ctx,
		container.ListOptions{
			Size:   true,
			All:    true,
			Latest: false,
			Limit:  c.pageConfig.SizeIntValue(),
		},
	)
}

func removePrefixSlash(s string) string {
	if strings.HasPrefix(s, "/") {
		return s[1:]
	}
	return s
}
