export VERSION=0.0.1
CMD_DIR=cmd
RELEASE_ROOT=release/nerv
PKG_AGENT=agent.tgz

build : cli server file resources  agent log elasticsearch config
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

log :
	@echo "----build log----"
	cd $(CMD_DIR)/log && make

elasticsearch :
	@echo "----build elasticsearch----"
	cd $(CMD_DIR)/elasticsearch && make

.PHONY : store
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
	mkdir $(RELEASE_ROOT)/resources/config
	cp -R profile/$(ENV)/ $(RELEASE_ROOT)
	cp -R profile/$(ENV)/ $(RELEASE_ROOT)/resources/config/nerv
	@echo "----config complete----"

.PHONY : pkg-service
pkg-service :
	@echo "----build pkg-service----"
	rm -rf release/nerv.tgz
	rm -rf $(RELEASE_ROOT)/pkg
	mkdir $(RELEASE_ROOT)/pkg
	mv $(RELEASE_ROOT)/agent.tgz $(RELEASE_ROOT)/pkg
	mv $(RELEASE_ROOT)/file.tgz $(RELEASE_ROOT)/pkg
	mv $(RELEASE_ROOT)/server.tgz $(RELEASE_ROOT)/pkg
	mv $(RELEASE_ROOT)/nerv-cli.tgz $(RELEASE_ROOT)/pkg
	mv $(RELEASE_ROOT)/log.tgz $(RELEASE_ROOT)/pkg
	mv $(RELEASE_ROOT)/elasticsearch.tgz $(RELEASE_ROOT)/pkg
	cp -R pkg/ $(RELEASE_ROOT)/pkg
	rm -rf $(RELEASE_ROOT)/file
	rm -rf $(RELEASE_ROOT)/server
	rm -rf $(RELEASE_ROOT)/agent
	rm -rf $(RELEASE_ROOT)/log
	rm -rf $(RELEASE_ROOT)/elasticsearch
	@echo "----pkg-service complete----"

.PHONY : pkg-webui
pkg-webui : webui
	@echo "----build pkg-webui----"
	mv $(RELEASE_ROOT)/webui.tgz $(RELEASE_ROOT)/pkg
	rm -rf $(RELEASE_ROOT)/webui
	@echo "----pkg-webui complete----"

.PHONY : pkg-all
pkg-all :
	@echo "----build pkg-all----"
	cd release && tar -zcvf nerv.tgz nerv
	@echo "----build pkg-all----"

.PHONY : resources
resources :
	@echo "----build resources----"
	cd resources && make
