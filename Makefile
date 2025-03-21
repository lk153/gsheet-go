GOPATH              := $(or $(GOPATH), $(HOME)/go)
GOLINT              := golangci-lint run
GOLINTCLEARCACHE	:= golangci-lint cache clean
GO_TEST_PARALLEL    := go test -parallel 4 -count=1 -timeout 30s
GOOGLE_WIRE 		:= $(GOPATH)/bin/wire
GOBUILDDEBUG        := go build -gcflags=all="-N -l"

$(MOCKERY):
	GOPATH=$(GOPATH) go install -mod=mod github.com/vektra/mockery/v2@latest
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
	$(GO_TEST_PARALLEL) ./... -v -coverprofile=cover.out && go tool cover -html=cover.out
generate:
	go generate ./...
start-debug: build-debug
	./gsheet-go