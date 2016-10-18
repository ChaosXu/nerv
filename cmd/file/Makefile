LDFLAGS=-ldflags "-X main.Version=$(VERSION)"

APP_SCRIPT=app
MODULE=file
RELEASE_DIR=../../release/nerv/$(MODULE)
BIN_DIR=$(RELEASE_DIR)/bin
CONF_DIR=$(RELEASE_DIR)/config
SCRIPTS_DIR=$(RELEASE_DIR)/scripts


file : clean
	go build $(LDFLAGS) -o $(BIN_DIR)/$(MODULE)
	cp bin/$(APP_SCRIPT) $(BIN_DIR)/$(APP_SCRIPT)
	chmod +x $(BIN_DIR)/$(APP_SCRIPT)
	mkdir $(CONF_DIR)
	cp config/config.json $(CONF_DIR)/config.json
	mkdir $(SCRIPTS_DIR)
	cp -R scripts/ $(SCRIPTS_DIR)
	
	
.PHONY : clean
clean :
	rm -rf $(RELEASE_DIR)