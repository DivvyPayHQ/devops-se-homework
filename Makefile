GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
PROTO_SRC_DIR=./proto

all: protoc build

protoc:
	protoc -I=$(PROTO_SRC_DIR) --go_out=plugins=grpc:./headquarters --go_out=plugins=grpc:./ship $(PROTO_SRC_DIR)/*.proto
build:
	$(GOBUILD) -o ship -v ./ship
	$(GOBUILD) -o headquarters -v ./headquarters
clean:
	rm -f ./ship/ship
	rm -f ./headquarters/headquarters