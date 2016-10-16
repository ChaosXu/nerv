LDFLAGS=-ldflags "-X main.Version=$(VERSION)"

RELEASE_DIR=../../release/nerv/server
BIN_DIR=$(RELEASE_DIR)/bin
CONF_DIR=$(RELEASE_DIR)/config

server : clean
	mkdir $(RELEASE_DIR)
	mkdir $(CONF_DIR)
	cp config/config.json $(CONF_DIR)/config.json 
	go build $(LDFLAGS) -o $(BIN_DIR)/server


.PHONY : clean
clean :
	rm -rf $(RELEASE_DIR)