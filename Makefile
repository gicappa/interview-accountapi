# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=client_example
BINARY_UNIX=$(BINARY_NAME)_unix

all: test build

build: 
	$(GOBUILD) -o $(BINARY_NAME) -v github.com/gicappa/interview-accountapi/cmd/client_example

test: 
	$(GOTEST) -v ./...

clean: 
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)

run:
	$(GOBUILD) -o $(BINARY_NAME) -v ./...
	./$(BINARY_NAME)

# Cross compilation
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v github.com/gicappa/interview-accountapi/cmd/client_example
	
docker-build:
	docker run --rm -it -v "$(GOPATH)":/go -w github.com/gicappa/interview-accountapi/cmd/client_example golang:latest go build -o "$(BINARY_UNIX)" -v github.com/gicappa/interview-accountapi/cmd/client_example

	# echo "Compiling for every OS and Platform"
	# GOOS=linux GOARCH=arm go build -o bin/main-linux-arm main.go
	# GOOS=linux GOARCH=arm64 go build -o bin/main-linux-arm64 main.go
	# GOOS=freebsd GOARCH=386 go build -o bin/main-freebsd-386 main.go