# Go parameters
.PHONY:  lint run depgraph build all test testv coverage threshold testl testall
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GORUN=$(GOCMD) run .
GOTEST=$(GOCMD) test -tags test -short
GOGET=$(GOCMD) get
GOCOV=$(GOCMD) tool cover -html=coverage.out
GODEP=godepgraph -s -o  github.com/webmalc/contacts-extractor-backend github.com/webmalc/contacts-extractor | dot -Tpng -o godepgraph.png
BINARY_NAME=contacts_extractor.app

threshold:
	overcover --coverprofile coverage.out --threshold 55 --summary

all: build

test:
	GOENV=test $(GOTEST) ./... -coverprofile=coverage.out

testv:
	GOENV=test $(GOTEST) -v ./... -coverprofile=coverage.ou

testl: testv lint

testall: test lint threshold

coverage:
	$(GOCOV)

build:
	$(GOBUILD) -o $(BINARY_NAME) -v

depgraph:
	$(GODEP)

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

lint:
	golangci-lint run ./...
run:
	$(GORUN) $(c)