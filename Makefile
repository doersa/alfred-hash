SHELL=/bin/bash

GO_MACOS_X86_PREFIX=CGO_ENABLED=0 GOOS=darwin GOARCH=amd64
GO_MACOS_ARM_PREFIX=CGO_ENABLED=0 GOOS=darwin GOARCH=arm64
GO_CMD=go
GO_BUILD=$(GO_CMD) build
GO_BUILD_COMMON=$(GO_BUILD)
GO_BUILD_MACOS_X86=$(GO_MACOS_X86_PREFIX) $(GO_BUILD)
GO_BUILD_MACOS_ARM=$(GO_MACOS_ARM_PREFIX) $(GO_BUILD)

.PHONY: clean all archive x86_app arm_app build
.DEFAULT: all

all: build 

x86_app:
	$(GO_MACOS_X86_PREFIX) $(GO_BUILD) -o x86_app

arm_app:
	$(GO_MACOS_ARM_PREFIX) $(GO_BUILD) -o arm_app

build: x86_app arm_app
	lipo -create -output hash x86_app arm_app; \
	rm -rf x86_app arm_app

unzip:
	unzip Hash.alfredworkflow -d tmp; 

archive: 
	zip -r Hash.alfredworkflow hash icon.png info.plist; 

clean:
	rm -rf arm_app hash x86_app \
	rm -rf tmp;
