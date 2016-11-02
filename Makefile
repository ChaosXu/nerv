export VERSION=0.0.1
CMD_DIR=cmd
RELEASE_ROOT=release/nerv
PKG_AGENT=agent.tar.gz

all : build
	@echo "----package----"
	rm -rf release/nerv.tar.gz
	cd release && tar -zcvf nerv.tar.gz nerv
	@echo "----package complete----"


build : server file  store pkg resources bin
	@echo "----build complete----"


server :
	@echo "----build server----"
	cd $(CMD_DIR)/server && make

file :
	@echo "----build file----"
	cd $(CMD_DIR)/file && make

agent :
	@echo "----build agent----"
	cd $(CMD_DIR)/agent && make

pkg : agent
	@echo "----build pkg----"
	rm -rf $(RELEASE_ROOT)/pkg
	mkdir $(RELEASE_ROOT)/pkg
	mv $(RELEASE_ROOT)/agent.tar.gz $(RELEASE_ROOT)/pkg

store :
	@echo "----build store----"
	cd $(CMD_DIR)/store && make

profile : all
	@echo "----profile $(ENV)----"
	cp -R profile/$(ENV)/ $(RELEASE_ROOT)
	@echo "----profile complete----"

.PHONY : resources
resources :
	@echo "----build resources----"
	cd resources && make

.PHONY : bin
bin :
	@echo "----build bin----"
	cd bin && make
