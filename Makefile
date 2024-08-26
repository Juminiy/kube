WORK_DIR:= $(PWD)
CMD_DIR := $(WORK_DIR)/cmd
BIN_DIR := $(WORK_DIR)/bin
LDFLAGS := $(shell version/version.sh)
GO_ENVS := CGO_ENABLED=0 GOOS=linux GOARCH=amd64
GO_BUILD:= $(GO_ENV) go build -ldflags "$(LDFLAGS)" -o $(BIN_DIR)/$@
GO_RUN_BUILD ?= cd $(CMD_DIR)/$@ && $(GO_BUILD)

.PHONY: set
set:
	mkdir -p $(BIN_DIR)
	GOPROXY=https://goproxy.cn,direct go mod tidy

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

.PHONY: clean
clean:
	rm -rf $(BIN_DIR)
