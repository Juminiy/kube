WORK_DIR:= $(PWD)
CMD_DIR := $(WORK_DIR)/cmd
BIN_DIR := $(WORK_DIR)/bin
LDFLAGS := $(shell version/version.sh)
GO_ENVS  = CGO_ENABLED=0

HOST_OS   := $(shell uname -s)
HOST_ARCH := $(shell uname -m)

ifeq ($(HOST_OS),Linux)
	GO_ENVS += GOOS=linux
endif
ifeq ($(HOST_OS),Darwin)
	GO_ENVS += GOOS=darwin
endif
ifeq ($(findstring NT,$(HOST_OS)),NT)
	GO_ENVS += GOOS=windows
endif

ifeq ($(HOST_ARCH),x86_64)
	GO_ENVS += GOARCH=amd64
endif
ifeq ($(HOST_ARCH),i386)
	GO_ENVS += GOARCH=386
endif
ifeq ($(HOST_ARCH),aarch64)
	GO_ENVS += GOARCH=arm64
endif
ifeq ($(HOST_ARCH),arm)
	GO_ENVS += GOARCH=arm
endif

GO_REMOVE_SYMBOL_TABLE := -s
GO_REMOVE_DEBUG_INFO   := -w
GO_BUILD	 := $(GO_ENVS) go build -ldflags "$(LDFLAGS) $(GO_REMOVE_SYMBOL_TABLE) $(GO_REMOVE_DEBUG_INFO)" -o $(BIN_DIR)/$@
GO_RUN_BUILD ?= cd $(CMD_DIR)/$@ && $(GO_BUILD)
GO_LIST_DIR   = `go list ./...`


.PHONY: set
set:
	mkdir -p $(BIN_DIR)
	env GO111MODULE=on GOPROXY=https://goproxy.cn,direct go mod tidy

.PHONY: all
all: menud consoled marketd payd
	@echo "make $^ in $(HOST_ARCH)/$(HOST_OS)"
	@echo $^

.PHONY: test
test:
	go test $(GO_LIST_DIR)
	go test -cover $(GO_LIST_DIR)

vet:
	go vet $(GO_LIST_DIR)

.PHONY: clean
clean:
	rm -rf $(BIN_DIR)

.PHONY: menud
menud: set vet
	$(GO_RUN_BUILD)

.PHONY: consoled
consoled: set vet
	$(GO_RUN_BUILD)

.PHONY: marketd
marketd: set vet
	$(GO_RUN_BUILD)

.PHONY: payd
payd: set vet
	$(GO_RUN_BUILD)