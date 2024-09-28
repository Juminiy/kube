package docker_internal

// please notice:
// lower case word, like: true, false, running, is constant
// UPPER CASE word, like: IMAGE_ID, IMAGE_NAME, is variable
// sign '|' is or
// sign '<>' is variable
// sign '[]' is optional

// filter args in image
// +referred from:
// https://docs.docker.com/reference/cli/docker/image/ls/#filter
// cli: docker image --filter "foo=bar" --filter "bif=baz"
const (
	// FilterDangling
	// +desc
	// image that without label
	// +usage
	// dangling=true|false
	FilterDangling = "dangling"

	// FilterLabel
	// +usage
	// label=com.example.version=1.0
	FilterLabel = "label"

	// FilterBefore
	// +usage
	// before=IMAGE_ID|IMAGE_NAME
	FilterBefore = "before"

	// FilterSince
	// +usage
	// before=IMAGE_ID|IMAGE_NAME
	FilterSince = "since"

	// FilterReference
	// reference=IMAGE_NAME:IMAGE_TAG
	FilterReference = "reference"
)

// filter args in container
// +referred from:
// https://docs.docker.com/reference/cli/docker/container/ls/#filter
// cli: docker ps --filter "foo=bar" --filter "bif=baz"
const (
	FilterID    = "id"
	FilterName  = "name"
	filterLabel = FilterLabel

	// FilterExited
	// +usage
	// exited=created|restarting|running|removing|paused|exited|dead
	FilterExited   = "exited"
	FilterStatus   = "status"
	FilterAncestor = "ancestor"
	filterBefore   = FilterBefore
	filterSince    = FilterSince
	FilterVolume   = "volume"
	FilterNetwork  = "network"

	// FilterPublish
	// FilterExpose
	// +usage
	// publish=PORT/[PROTO] | STARTPORT-ENDPORT/[PROTO]
	FilterPublish = "publish"
	FilterExpose  = "expose"

	// FilterHealth
	// +usage
	// health=starting|healthy|unhealthy|none
	FilterHealth = "health"

	// FilterIsolation
	// +usage
	// isolation=default|process|hyperv
	FilterIsolation = "isolation"

	// FilterIsTask
	// +usage
	// is-task=true|false
	FilterIsTask = "is-task"
)

// filter constant
// +self define
const (
	// referred from: github.com/docker/docker/client/image_list_test.go/TestImageListConnectionError
	ReferenceNone = "no-such-image.invalid:no-such-tag.invalid"
	ReferenceAll  = "*"
)
