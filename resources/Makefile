MODULE=resources
RELEASE_ROOT=../release/nerv
RELEASE_DIR=$(RELEASE_ROOT)/$(MODULE)

resources : clean	
	mkdir $(RELEASE_DIR)
	find . -name '*.sh' | xargs tar -czf $(MODULE).tar.gz
	tar -xzf $(MODULE).tar.gz -C $(RELEASE_DIR)
	find $(RELEASE_DIR) -name '*.sh' | xargs chmod +x
	rm -rf $(MODULE).tar.gz
	find . -name '*.json' | xargs tar -czf $(MODULE).tar.gz
	tar -xzf $(MODULE).tar.gz -C $(RELEASE_DIR)
	rm -rf $(MODULE).tar.gz
	
.PHONY : clean
clean :
	rm -rf $(RELEASE_DIR)