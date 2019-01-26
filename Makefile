SOURCE_FILES := Makefile $(shell find . -name '*.go')
VERSION_LDFLAGS=-X main.version=$(shell git describe --always --long --dirty)

BINARY_PATH=bin
BINARY_LINUX=kubenv-linux-amd64
BINARY_DARWIN=kubenv-darwin-amd64
BINARY_WINDOWS=kubenv-windows-amd64.exe

BIN_PATH=$(shell go env GOPATH)/bin

default: build



# Build
build: $(BINARY_PATH)/$(BINARY_LINUX) $(BINARY_PATH)/$(BINARY_DARWIN) $(BINARY_PATH)/$(BINARY_WINDOWS)
	
$(BINARY_PATH)/$(BINARY_LINUX): $(SOURCE_FILES)
	GOARCH=amd64 GOOS=linux go build -i -v -ldflags="$(VERSION_LDFLAGS)" -o $(BINARY_PATH)/$(BINARY_LINUX)

$(BINARY_PATH)/$(BINARY_DARWIN): $(SOURCE_FILES)
	GOARCH=amd64 GOOS=darwin go build -i -v -ldflags="$(VERSION_LDFLAGS)" -o $(BINARY_PATH)/$(BINARY_DARWIN)

$(BINARY_PATH)/$(BINARY_WINDOWS): $(SOURCE_FILES)
	GOARCH=amd64 GOOS=windows go build -i -v -ldflags="$(VERSION_LDFLAGS)" -o $(BINARY_PATH)/$(BINARY_WINDOWS)



# Clean
clean: clean-$(BINARY_PATH)/$(BINARY_LINUX) clean-$(BINARY_PATH)/$(BINARY_DARWIN) clean-$(BINARY_PATH)/$(BINARY_WINDOWS)

clean-$(BINARY_PATH)/$(BINARY_LINUX):
	rm -f $(BINARY_PATH)/$(BINARY_LINUX)

clean-$(BINARY_PATH)/$(BINARY_DARWIN):
	rm -f $(BINARY_PATH)/$(BINARY_DARWIN)

clean-$(BINARY_PATH)/$(BINARY_WINDOWS):
	rm -f $(BINARY_PATH)/$(BINARY_WINDOWS)



# Install

install: install-$(BINARY_PATH)/$(BINARY_DARWIN)

install-$(BINARY_PATH)/$(BINARY_LINUX): $(BINARY_PATH)/$(BINARY_LINUX)
	cp $(BINARY_PATH)/$(BINARY_LINUX) $(BIN_PATH)/kubenv

install-$(BINARY_PATH)/$(BINARY_DARWIN): $(BINARY_PATH)/$(BINARY_DARWIN)
	cp $(BINARY_PATH)/$(BINARY_DARWIN) $(BIN_PATH)/kubenv

install-$(BINARY_PATH)/$(BINARY_WINDOWS): $(BINARY_PATH)/$(BINARY_WINDOWS)
	cp $(BINARY_PATH)/$(BINARY_WINDOWS) $(BIN_PATH)/kubenv
