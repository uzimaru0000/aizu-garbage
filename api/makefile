GOCMD=go
GOBUILD=$(GOCMD) build
GOGET=$(GOCMD) get
BINARY_NAME=api
BINARY_UNIX=$(BINARY_NAME)_unix

all: setup build
build:
	$(GOBUILD) -o $(BINARY_NAME)
setup:
	$(GOGET) -u github.com/golang/dep/cmd/dep
	dep ensure