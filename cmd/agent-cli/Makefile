LDFLAGS=-ldflags "-X main.Version=$(VERSION)"

APP_SCRIPT=app
MODULE=agent-cli
PKG=$(MODULE)-$(PLATFORM).tgz
RELEASE_ROOT=../../release/nerv
RELEASE_DIR=$(RELEASE_ROOT)/agent
BIN_DIR=$(RELEASE_DIR)/bin
CONF_DIR=$(RELEASE_DIR)/config


agent :
	go build $(LDFLAGS) -o $(BIN_DIR)/$(MODULE)

