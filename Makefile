export VERSION=0.0.1
CMD_DIR=cmd
RELEASE_ROOT=release/nerv
PKG_AGENT=agent.tar.gz

build : cli server file  store resources  agent config
	@echo "----build complete----"

cli :
	@echo "----build cli----"
	cd $(CMD_DIR)/cli && make

server :
	@echo "----build server----"
	cd $(CMD_DIR)/server && make

file :
	@echo "----build file----"
	cd $(CMD_DIR)/file && make

agent :
	@echo "----build agent----"
	cd $(CMD_DIR)/agent && make

store :
	@echo "----build store----"
	cd $(CMD_DIR)/store && make

.PHONY : webui
webui :
	@echo "----build webui----"
	cd $(CMD_DIR)/webui && make


.PHONY : config
config :
	@echo "----config $(ENV)----"
	cp -R profile/$(ENV)/ $(RELEASE_ROOT)
	@echo "----config complete----"

.PHONY : pkg
pkg :
	@echo "----build pkg----"
	rm -rf release/nerv.tar.gz
	rm -rf $(RELEASE_ROOT)/pkg
	mkdir $(RELEASE_ROOT)/pkg
	mv $(RELEASE_ROOT)/agent.tar.gz $(RELEASE_ROOT)/pkg
	mv $(RELEASE_ROOT)/file.tar.gz $(RELEASE_ROOT)/pkg
	mv $(RELEASE_ROOT)/server.tar.gz $(RELEASE_ROOT)/pkg
	mv $(RELEASE_ROOT)/nerv-cli.tar.gz $(RELEASE_ROOT)/pkg
	mv $(RELEASE_ROOT)/webui.tar.gz $(RELEASE_ROOT)/pkg
	rm -rf $(RELEASE_ROOT)/file
	rm -rf $(RELEASE_ROOT)/server
	rm -rf $(RELEASE_ROOT)/agent
	rm -rf $(RELEASE_ROOT)/webui
	cd release && tar -zcvf nerv.tar.gz nerv
	@echo "----pkg complete----"


.PHONY : resources
resources :
	@echo "----build resources----"
	cd resources && make
