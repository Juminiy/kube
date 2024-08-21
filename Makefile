BIN_NAM := kube
BIN_DIR := bin
LDFLAGS := $(shell version/version.sh)

.PHONY: set
set:
	rm -rf $(BIN_DIR)
	mkdir -p $(BIN_DIR)
	GOPROXY=https://goproxy.cn,direct go mod tidy

.PHONY: kube
kube: set
	CGO_ENABLED=0 go build -ldflags "$(LDFLAGS)" -o $(BIN_DIR)/$(BIN_NAM)