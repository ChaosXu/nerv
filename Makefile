CMD_DIR=cmd
all : server
	@echo "----build complete----"

.PHONY : server 
server :
	@echo "----build server----"
	cd $(CMD_DIR)/server && make server
