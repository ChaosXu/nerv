MODULE=scripts
RELEASE_ROOT=../release/nerv
RELEASE_DIR=$(RELEASE_ROOT)/$(MODULE)

scripts : clean	
	mkdir $(RELEASE_DIR)
	find . -name '*.sh' | xargs tar -czf $(MODULE).tar.gz
	tar -xzf $(MODULE).tar.gz -C $(RELEASE_DIR)
	rm -rf $(MODULE).tar.gz
	
.PHONY : clean
clean :
	rm -rf $(RELEASE_DIR)