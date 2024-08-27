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

GO_BUILD	 := $(GO_ENVS) go build -ldflags "$(LDFLAGS)" -o $(BIN_DIR)/$@
GO_RUN_BUILD ?= cd $(CMD_DIR)/$@ && $(GO_BUILD)


.PHONY: set
set:
	mkdir -p $(BIN_DIR)
	GOPROXY=https://goproxy.cn,direct go mod tidy
	# fill cmd

.PHONY: menud
menud: set
	$(GO_RUN_BUILD)

.PHONY: imaged
imaged: set
	$(GO_RUN_BUILD)

.PHONY: instanced
instanced: set
	$(GO_RUN_BUILD)

.PHONY: marketd
marketd: set
	$(GO_RUN_BUILD)

.PHONY: all
all: menud imaged instanced marketd
	@echo "make all in $(HOST_ARCH)/$(HOST_OS)"
	@echo $^

.PHONY: clean
clean:
	rm -rf $(BIN_DIR)