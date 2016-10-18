export VERSION=0.0.1
CMD_DIR=cmd
RELEASE_ROOT=release/nerv
PKG_AGENT=agent.tar.gz

all : build
	

build : server agent file
	cp $(RELEASE_ROOT)/$(PKG_AGENT) $(RELEASE_ROOT)/file/pkg 
	@echo "----build complete----"

 
server :
	@echo "----build server----"
	cd $(CMD_DIR)/server && make

agent :
	@echo "----build agent----"
	cd $(CMD_DIR)/agent && make

file :
	@echo "----build file----"
	cd $(CMD_DIR)/file && make
