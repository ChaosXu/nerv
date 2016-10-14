RELEASE_DIR=../../release/server
BIN_DIR=${RELEASE_DIR}/bin
CONF_DIR=${RELEASE_DIR}/conf

server : clean
	go build -o $(BIN_DIR)/server


.PHONY : clean
clean :
	rm -rf $(RELEASE_DIR)