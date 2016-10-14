RELEASE_DIR=../../release/server
BIN_DIR=${RELEASE_DIR}/bin
CONF_DIR=${RELEASE_DIR}/conf

clean :
	rm -rf $(RELEASE_DIR)

build : clean 
	go build -o $(BIN_DIR)/server
