GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
BINARY_NAME=api
GCLOUD_PROJECT_ID=aizu-garbage-260900

all: build
build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_NAME) -v
deploy:
	gcloud builds submit --project $(GCLOUD_PROJECT_ID) --tag gcr.io/$(GCLOUD_PROJECT_ID)/api-image
