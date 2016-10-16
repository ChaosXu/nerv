export VERSION=0.0.1
CMD_DIR=cmd

all : server
	@echo "----tar----"
	rm -rf release/nerv.tar.gz
	tar -zcvf release/nerv.tar.gz release/nerv 
	@echo "----build complete----"

.PHONY : server 
server :
	@echo "----build server----"
	cd $(CMD_DIR)/server && make server
