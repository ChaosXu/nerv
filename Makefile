export VERSION=0.0.1
CMD_DIR=cmd

all : build
	@echo "----tar----"
	rm -rf release/nerv.tar.gz
	cd release
	tar -zcvf nerv.tar.gz nerv 
	@echo "----tar complete----"

build : server
	@echo "----build complete----"

.PHONY : server 
server :
	@echo "----build server----"
	cd $(CMD_DIR)/server && make server
