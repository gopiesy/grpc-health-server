GO_PKG_DIRS  := $(subst $(shell go list -e -m),.,$(shell go list ./... | grep -v /vendor | grep -v /health-server ))

all: clean fmt lint
	go build -ldflags="-s -w" -o server $(GO_PKG_DIRS)

fmt:
	gofmt -s -w $(GO_PKG_DIRS)

lint:
	golangci-lint run -v $(GO_PKG_DIRS)

proto:
	./proto/run

clean:
	rm -f server
