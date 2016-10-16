LDFLAGS=-ldflags "-X main.Version=$(VERSION)"

RELEASE_DIR=../../release/nerv/server
BIN_DIR=$(RELEASE_DIR)/bin
CONF_DIR=$(RELEASE_DIR)/config

server : clean
	go build $(LDFLAGS) -o $(BIN_DIR)/server
	mkdir $(CONF_DIR)
	cp config/config.json $(CONF_DIR)/config.json

.PHONY : clean
clean :
	rm -rf $(RELEASE_DIR)