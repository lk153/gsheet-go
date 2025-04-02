.PHONY: data

GOPATH              := $(or $(GOPATH), $(HOME)/go)
GOLINT              := GO111MODULE=on CGO_ENABLED=1 golangci-lint run
GOLINTCLEARCACHE	:= GO111MODULE=on CGO_ENABLED=1 golangci-lint cache clean
GO_TEST_PARALLEL    := go test
GOOGLE_WIRE 		:= $(GOPATH)/bin/wire
MOCKERY 			:= $(GOPATH)/bin/mockery
GOBUILDDEBUG        := go build -gcflags=all="-N -l"
SOURCE=$(shell go list ./... | grep -v /mocks/)

$(MOCKERY):
	GOPATH=$(GOPATH) go install github.com/vektra/mockery/v2@latest
$(GOOGLE_WIRE):
	GOPATH=$(GOPATH) go install github.com/google/wire/cmd/wire@latest

clean:
	rm -rf gsheet-go cpu.pprof mem.pprof
build: $(GOOGLE_WIRE) clean
	go mod tidy && go mod vendor
build-debug: clean
	$(GOBUILDDEBUG)
lint:
	$(GOLINTCLEARCACHE) && $(GOLINT) -v ./...
test:
	$(GO_TEST_PARALLEL) $(SOURCE) -v -coverprofile=cover.out && go tool cover -html=cover.out
generate:
	go generate ./...
start-debug: build-debug
	./gsheet-go
mock: $(MOCKERY)
	rm -rf mocks && $(MOCKERY)