# Makefile for building and running the application

```makefile
# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod
BINARY_NAME=contained
DOCKER_IMAGE_NAME=contained

# Main package path
MAIN_PACKAGE_PATH=./cmd/contained/main.go

all: test build

build:
	$(GOBUILD) -o $(BINARY_NAME) -v $(MAIN_PACKAGE_PATH)

test:
	$(GOTEST) -v ./...

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

run:
	$(GOBUILD) -o $(BINARY_NAME) -v $(MAIN_PACKAGE_PATH)
	./$(BINARY_NAME)

deps:
	$(GOGET) ./...
	$(GOMOD) tidy

docker-build:
	docker build -t $(DOCKER_IMAGE_NAME) -f .docker/Dockerfile .

docker-run:
	docker run -p 8080:8080 $(DOCKER_IMAGE_NAME)

.PHONY: all build test clean run deps docker-build docker-run
```