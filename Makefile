# Force-enable Go modules even if this project has been cloned within a user's GOPATH
export GO111MODULE = on

# Specify VERBOSE=1 to get verbose output from all executed commands
ifdef VERBOSE
V = -v
X = -x
else
.SILENT:
endif

.PHONY: all
all: build test

.PHONY: clean
clean:
	rm -rf bin/ coverage/ cucumber/logs/
	go clean -i $(X) -cache -testcache

.PHONY: build
build:
	mkdir -p bin
	go build $(V) -o bin/gokay

.PHONY: fmt
fmt:
	go fmt $(X) ./...

.PHONY: test
test:
	mkdir -p coverage
	go test $(V) -race -cover -coverprofile coverage/cover.profile ./...

.PHONY: cover
cover:
	go tool cover -html coverage/cover.profile
