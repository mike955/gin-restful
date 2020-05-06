#=============================================================================
#     FileName: Makefile
#         Desc: utils's Makefile 
#       Author: ato.ye
#        Email: ato.ye@ucloud.cn
#     HomePage: http://www.ucloud.cn
#      Version: 0.0.1
#   LastChange: 2016-04-20 18:18:19
#      History:
#=============================================================================

SRCDIRS_EXCLUDE = new_project zk_admin 
SRCDIRS_ALL = $(sort $(subst /,,$(dir $(wildcard */))))
SRCDIRS = $(filter-out $(SRCDIRS_EXCLUDE), $(SRCDIRS_ALL))

PKGDIRS_EXCLUDE = $(GOROOT)/pkg
PKGDIRS_ALL = $(addsuffix /pkg, $(subst :, ,$(GOPATH)))
PKGDIRS = $(filter-out $(PKGDIRS_EXCLUDE), $(PKGDIRS_ALL))


all:zkadmin
	@for subdir in $(SRCDIRS);do \
		cd $$subdir; go install; cd $(CURDIR); \
	done 

zkadmin:
	@cd zk_admin; \
	CGO_ENABLED=0 go install;\
	cd $(CURDIR);

show:
	@echo SRCDIRS: $(SRCDIRS)
	@echo PKGDIRS: $(PKGDIRS)

clean: 
	@go clean -i ./...
