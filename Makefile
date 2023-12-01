GOPATH              := $(or $(GOPATH), $(HOME)/go)
GOLINT              := golangci-lint run
GOLINTCLEARCACHE	:= golangci-lint cache clean
GO_TEST_PARALLEL    := go test -parallel 4 -count=1 -timeout 30s
GOSTATIC            := go build -ldflags="-w -s"

clean:
	rm -rf ./out/main cpu.pprof mem.pprof
build: clean
	go mod tidy && go mod vendor && $(GOSTATIC) -o out/main ./
lint:
	$(GOLINTCLEARCACHE) && $(GOLINT) -v ./...
test:
	$(GO_TEST_PARALLEL) ./... -v -coverprofile=cover.out && go tool cover -html=cover.out
generate:
	go generate ./...