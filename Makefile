WORK_DIR:= $(PWD)
CMD_DIR := $(WORK_DIR)/cmd
BIN_DIR := $(WORK_DIR)/bin
TAR_DIR  = $(BIN_DIR)
LDFLAGS := $(shell hack/version.sh)
GO_ENVS  = env CGO_ENABLED=0

# default host arch/os as target
HOST_ARCH := $(shell uname -m)
HOST_OS   := $(shell uname -s)

ifeq ($(HOST_ARCH),x86_64)
	GO_ENVS += GOARCH=amd64
endif
ifeq ($(HOST_ARCH),i386)
	GO_ENVS += GOARCH=386
endif
ifeq ($(HOST_ARCH),aarch64)
	GO_ENVS += GOARCH=arm64
endif
ifeq ($(HOST_ARCH),arm64)
	GO_ENVS += GOARCH=arm64
endif
ifeq ($(HOST_ARCH),arm)
	GO_ENVS += GOARCH=arm
endif

ifeq ($(HOST_OS),Linux)
	GO_ENVS += GOOS=linux
endif
ifeq ($(HOST_OS),Darwin)
	GO_ENVS += GOOS=darwin
endif
ifeq ($(findstring NT,$(HOST_OS)),NT)
	GO_ENVS += GOOS=windows
endif

# go compile cmd
GO_RM_SYMBOL_TABLE := -s
GO_RM_DEBUG_INFO   := -w
GO_BUILD	 ?= go build -ldflags "$(LDFLAGS) $(GO_RM_SYMBOL_TABLE) $(GO_RM_DEBUG_INFO)" -o $(TAR_DIR)/$@
GO_RUN_BUILD ?= cd $(CMD_DIR)/$@ && $(GO_ENVS) $(GO_BUILD)
GO_LIST_DIR   = `go list ./...`

# make target
.PHONY: set
set:
	mkdir -p $(BIN_DIR)
	env GO111MODULE=on GOPROXY=https://goproxy.cn,direct go mod tidy

.PHONY: test
test:
	go test $(GO_LIST_DIR)
	go test -cover $(GO_LIST_DIR)

.PHONY: vet
vet:
	go vet $(GO_LIST_DIR)

.PHONY: vendor
vendor:
	go mod vendor

.PHONY: vendortidy
vendortidy:
	rm -rf vendor


################################
#		codegen git-add  	   #
################################
.PHONY: codegen
codegen: set
	go run cmd/codegencli/codegencli.go -gen all
	git add pkg/image_api/docker_api/docker_inst/client.go
	git add pkg/image_api/harbor_api/harbor_inst/client.go
	git add pkg/storage_api/minio_api/minio_inst/client.go

.PHONY: codegendocker
codegendocker: set
	go run cmd/codegencli/codegencli.go -gen docker
	git add pkg/image_api/docker_api/docker_inst/client.go

.PHONY: codegenharbor
codegenharbor: set
	go run cmd/codegencli/codegencli.go -gen harbor
	git add pkg/image_api/harbor_api/harbor_inst/client.go

.PHONY: codegenminio
codegenminio: set
	go run cmd/codegencli/codegencli.go -gen minio
	git add pkg/storage_api/minio_api/minio_inst/client.go


################################
#		binary executable  	   #
################################
.PHONY: all
all: menud
	@echo "make $^ in $(HOST_ARCH)/$(HOST_OS)"

.PHONY: clean
clean:
	rm -rf $(BIN_DIR)

.PHONY: menud
menud: set vet
	$(GO_RUN_BUILD)

.PHONY: egvcheck
egvcheck: set vet
	$(GO_RUN_BUILD)

.PHONY: healthz
healthz: set vet
	$(GO_RUN_BUILD)

.PHONY: clipboard_fast
clipboard_fast:
	$(eval GO_ENVS = env CGO_ENABLED=1 GOOS=linux GOARCH=amd64)
	cd $(CMD_DIR)/clipboard/$@ && $(GO_ENVS) $(GO_BUILD)

.PHONY: hardcliv2
hardcliv2:
	$(eval GO_ENVS = env CGO_ENABLED=0 GOOS=linux GOARCH=amd64)
	$(GO_RUN_BUILD)

.PHONY: dctl
dctl: set vet
	$(eval GO_ENVS = env CGO_ENABLED=1)
	$(GO_RUN_BUILD)

.PHONY: dctlv2
dctlv2: set vet
	$(eval GO_ENVS = env CGO_ENABLED=1)
	$(GO_RUN_BUILD)

.PHONY: dctlv3
dctlv3: set vet
	$(eval GO_ENVS = env CGO_ENABLED=1)
	$(GO_RUN_BUILD)


################################
#		docker executable  	   #
################################
.PHONY: docker_clipboard_fast
docker_clipboard_fast:
	docker buildx build --platform linux/amd64,linux/arm64 \
		--builder multiple-x \
		--debug \
		--load \
		--tag clipboard_fast:latest \
		--file docker/clipboard_fast/Dockerfile .
