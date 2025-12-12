TOPDIR=$(shell pwd)

all: gendb
	@echo "ALL DONE"

gendb:
	@echo "start gen db model..."$(TOPDIR)
	@cd $(TOPDIR)/adaptor/repo/default && gentool -c gen.yaml
	@cd $(TOPDIR)/adaptor/repo/example && gentool -c gen.yaml