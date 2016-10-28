export VERSION=0.0.1
CMD_DIR=cmd
RELEASE_ROOT=release/nerv
PKG_AGENT=agent.tar.gz

all : build
	@echo "----package----"
	rm -rf release/nerv.tar.gz
	cd release && tar -zcvf nerv.tar.gz nerv
	@echo "----package complete----"
	

build : server agent file resoruces
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


profile : all	
	@echo "----profile $(ENV)----"
	cp -R profile/$(ENV)/ $(RELEASE_ROOT)
	@echo "----profile complete----"	

.PHONY : resources
resoruces :
	@echo "----build resources----"
	cd resources && make
