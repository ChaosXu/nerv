APP_SCRIPT=app
SETUP_SCRIPT=setup.sh
MODULE=store
RELEASE_ROOT=../../release/nerv
RELEASE_DIR=$(RELEASE_ROOT)/$(MODULE)
BIN_DIR=$(RELEASE_DIR)/bin


store : clean
	mkdir $(RELEASE_DIR)
	find . -name '*.sh' | xargs tar -czf $(MODULE).tar.gz
	tar -xzf $(MODULE).tar.gz -C $(RELEASE_DIR)
	find $(RELEASE_DIR) -name '*.sh' | xargs chmod +x
	cp bin/$(APP_SCRIPT) $(BIN_DIR)/$(APP_SCRIPT)
	chmod +x $(BIN_DIR)/$(APP_SCRIPT)


.PHONY : clean
clean :
	rm -rf $(RELEASE_DIR)
