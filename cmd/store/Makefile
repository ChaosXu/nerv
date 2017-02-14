APP_SCRIPT=app
MODULE=store
PKG=$(MODULE)-$(PLATFORM).tgz
RELEASE_ROOT=../../release/nerv
RELEASE_DIR=$(RELEASE_ROOT)/$(MODULE)
BIN_DIR=$(RELEASE_DIR)/bin

log : clean
	mkdir $(RELEASE_DIR)
	mkdir $(BIN_DIR)
	cp -R bin/ $(BIN_DIR)
	chmod +x $(BIN_DIR)/*
	rm -rf $(RELEASE_DIR).tgz
	cd $(RELEASE_ROOT) && tar -zcvf $(PKG) $(MODULE)

.PHONY : clean
clean :
	rm -rf $(RELEASE_DIR)
