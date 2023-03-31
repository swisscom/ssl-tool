BINARY_NAME=ssl-tool

build:
	CGO_ENABLED=0 go build -o "./build/$(BINARY_NAME)" ./

install:
	go install ./