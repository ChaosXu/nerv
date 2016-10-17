export VERSION=0.0.1
CMD_DIR=cmd

all : build
	@echo "----package----"
	rm -rf release/nerv.tar.gz
	cd release && tar -zcvf nerv.tar.gz nerv
	@echo "----package complete----"

build : server agent
	@echo "----build complete----"

 
server :
	@echo "----build server----"
	cd $(CMD_DIR)/server && make server

agent :
	@echo "----build agent----"
	cd $(CMD_DIR)/agent && make agent
