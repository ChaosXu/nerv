export VERSION=0.0.1
CMD_DIR=cmd
RELEASE_ROOT=release/nerv
PKG_AGENT=agent.tgz
OS=darwin
CPU=x86_64
PLATFORM=$(OS)-$(CPU)

build : cli server file agent log store logui resources config
	@echo "----build complete----"

cli :
	@echo "----build cli----"
	echo $(PLATFORM)
	cd $(CMD_DIR)/cli && make -e PLATFORM=$(PLATFORM)

server :
	@echo "----build server----"
	cd $(CMD_DIR)/server && make -e PLATFORM=$(PLATFORM)

file :
	@echo "----build file----"
	cd $(CMD_DIR)/file && make -e PLATFORM=$(PLATFORM)

agent :
	@echo "----build agent----"
	cd $(CMD_DIR)/agent && make -e PLATFORM=$(PLATFORM)

log :
	@echo "----build log----"
	cd $(CMD_DIR)/log && make -e PLATFORM=$(PLATFORM)

store :
	@echo "----build elasticsearch----"
	cd $(CMD_DIR)/store && make -e PLATFORM=$(PLATFORM)

logui :
	@echo "----build log----"
	cd $(CMD_DIR)/logui && make -e PLATFORM=$(PLATFORM)


webui :
	@echo "----build webui----"
	cd $(CMD_DIR)/webui && make -e PLATFORM=$(PLATFORM)

pkg-service :
	@echo "----build pkg-service----"
	rm -rf release/nerv-$(PLATFORM).tgz
	rm -rf $(RELEASE_ROOT)/pkg
	mkdir $(RELEASE_ROOT)/pkg
	mv $(RELEASE_ROOT)/agent-$(PLATFORM).tgz $(RELEASE_ROOT)/pkg
	mv $(RELEASE_ROOT)/file-$(PLATFORM).tgz $(RELEASE_ROOT)/pkg
	mv $(RELEASE_ROOT)/server-$(PLATFORM).tgz $(RELEASE_ROOT)/pkg
	mv $(RELEASE_ROOT)/nerv-cli-$(PLATFORM).tgz $(RELEASE_ROOT)/pkg
	mv $(RELEASE_ROOT)/log-$(PLATFORM).tgz $(RELEASE_ROOT)/pkg
	mv $(RELEASE_ROOT)/store-$(PLATFORM).tgz $(RELEASE_ROOT)/pkg
	mv $(RELEASE_ROOT)/logui-$(PLATFORM).tgz $(RELEASE_ROOT)/pkg
ifeq (pkg, $(wildcard pkg))
	cp -R pkg/ $(RELEASE_ROOT)/pkg
endif
	rm -rf $(RELEASE_ROOT)/file
	rm -rf $(RELEASE_ROOT)/server
	rm -rf $(RELEASE_ROOT)/agent
	rm -rf $(RELEASE_ROOT)/log
	rm -rf $(RELEASE_ROOT)/store
	rm -rf $(RELEASE_ROOT)/logui
	@echo "----pkg-service complete----"

.PHONY : pkg-webui
pkg-webui : webui
	@echo "----build pkg-webui----"
	mv $(RELEASE_ROOT)/webui-$(PLATFORM).tgz $(RELEASE_ROOT)/pkg
	rm -rf $(RELEASE_ROOT)/webui
	@echo "----pkg-webui complete----"


pkg-all :
	@echo "----build pkg-all----"
	cd release && tar -zcvf nerv-$(PLATFORM).tgz nerv
	@echo "----pkg-all complete----"

.PHONY : resources
resources :
	@echo "----build resources----"
	cd resources && make
	@echo "----resources complete----"

config :
	@echo "----config $(ENV)----"
	mkdir -p $(RELEASE_ROOT)/resources/config
	cp -R resources/config/ $(RELEASE_ROOT)/resources/config
	cp -R profile/$(ENV)/ $(RELEASE_ROOT)
	cp -R profile/$(ENV)/ $(RELEASE_ROOT)/resources/config/nerv
	@echo "----config complete----"
