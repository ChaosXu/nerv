LDFLAGS=-ldflags "-X main.Version=$(VERSION)"

APP_SCRIPT=app
MODULE=nerv-cli
PKG=$(MODULE).tar.gz
RELEASE_ROOT=../../release/nerv
RELEASE_DIR=$(RELEASE_ROOT)/$(MODULE)
BIN_DIR=$(RELEASE_DIR)/bin
CONF_DIR=$(RELEASE_DIR)/config


cli : clean
	go build $(LDFLAGS) -o $(BIN_DIR)/$(MODULE)
	rm -rf $(RELEASE_DIR).tar.gz
	cd $(RELEASE_ROOT) && tar -zcvf $(PKG) $(MODULE)


.PHONY : clean
clean :
	rm -rf $(RELEASE_DIR)
