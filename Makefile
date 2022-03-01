IMPORT_PATH := github.com/a_antonenko/help-ukraine
GO          := $(or $(shell command -v go 2> /dev/null),which go)
BUILD_DIR := $(CURDIR)/build

all: init build

env:    ## Print useful environment variables to stdout
	@echo CURDIR: $(CURDIR)
	@echo BUILD_DIR: $(BUILD_DIR)
	@echo GOROOT: $(GOROOT)
	@echo GO: $(GO)
	@echo GO_VERSION: $(GoVersion)

build: init
	mkdir -p ${BUILD_DIR}
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 $(GO) build ${FLAGS} -o ${BUILD_DIR}/darwin-amd64/help-ukraine
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 $(GO) build ${FLAGS} -o ${BUILD_DIR}/linux-amd64/help-ukraine
docker-build: init fmt vet
init:
	mkdir -p ${BUILD_DIR}
	$(GO) get -t .
fmt:
	$(GO) fmt $(IMPORT_PATH)...
vet:
	$(GO) vet -composites=false $(IMPORT_PATH)...
clean:
	rm -rf ${BUILD_DIR}
