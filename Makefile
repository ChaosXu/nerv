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

.PHONY : pkg-service
pkg-service :
	@echo "----build pkg-service----"
	rm -rf release/nerv.tar.gz
	rm -rf $(RELEASE_ROOT)/pkg
	mkdir $(RELEASE_ROOT)/pkg
	mv $(RELEASE_ROOT)/agent.tar.gz $(RELEASE_ROOT)/pkg
	mv $(RELEASE_ROOT)/file.tar.gz $(RELEASE_ROOT)/pkg
	mv $(RELEASE_ROOT)/server.tar.gz $(RELEASE_ROOT)/pkg
	mv $(RELEASE_ROOT)/nerv-cli.tar.gz $(RELEASE_ROOT)/pkg
	rm -rf $(RELEASE_ROOT)/file
	rm -rf $(RELEASE_ROOT)/server
	rm -rf $(RELEASE_ROOT)/agent
	@echo "----pkg-service complete----"

.PHONY : pkg-webui
pkg-webui : webui
	@echo "----build pkg-webui----"
	mv $(RELEASE_ROOT)/webui.tar.gz $(RELEASE_ROOT)/pkg
	rm -rf $(RELEASE_ROOT)/webui
	@echo "----pkg-webui complete----"

.PHONY : pkg-all
pkg-all :
	@echo "----build pkg-all----"
	cd release && tar -zcvf nerv.tar.gz nerv
	@echo "----build pkg-all----"

.PHONY : resources
resources :
	@echo "----build resources----"
	cd resources && make
