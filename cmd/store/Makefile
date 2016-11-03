APP_SCRIPT=app
SETUP_SCRIPT=setup.sh
MODULE=store
RELEASE_ROOT=../../release/nerv
RELEASE_DIR=$(RELEASE_ROOT)/$(MODULE)
BIN_DIR=$(RELEASE_DIR)/bin
ELASTICSEARCH=elasticsearch-5.0.0.tar.gz

store : clean
	mkdir $(RELEASE_DIR)
	find . -name '*.sh' | xargs tar -czf $(MODULE).tar.gz
	tar -xzf $(MODULE).tar.gz -C $(RELEASE_DIR)
	rm -f $(MODULE).tar.gz
	find $(RELEASE_DIR) -name '*.sh' | xargs chmod +x
	cp bin/$(APP_SCRIPT) $(BIN_DIR)/$(APP_SCRIPT)
	chmod +x $(BIN_DIR)/$(APP_SCRIPT)
ifeq ($(ENV) , debug)
	cp bin/$(ELASTICSEARCH)  $(BIN_DIR)/$(ELASTICSEARCH)
endif

.PHONY : clean
clean :
	rm -rf $(RELEASE_DIR)
