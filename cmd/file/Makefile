LDFLAGS=-ldflags "-X main.Version=$(VERSION)"

APP_SCRIPT=app
MODULE=file
PKG=$(MODULE)-$(PLATFORM).tgz
RELEASE_ROOT=../../release/nerv
RELEASE_DIR=$(RELEASE_ROOT)/$(MODULE)
BIN_DIR=$(RELEASE_DIR)/bin
CONF_DIR=$(RELEASE_DIR)/config


server : clean
	go build $(LDFLAGS) -o $(BIN_DIR)/$(MODULE)
	cp bin/$(APP_SCRIPT) $(BIN_DIR)/$(APP_SCRIPT)
	chmod +x $(BIN_DIR)/$(APP_SCRIPT)
	mkdir $(CONF_DIR)
	cp config/config.json $(CONF_DIR)/config.json
	rm -rf $(RELEASE_DIR).tgz
	cd $(RELEASE_ROOT) && tar -zcvf $(PKG) $(MODULE)


.PHONY : clean
clean :
	rm -rf $(RELEASE_DIR)
