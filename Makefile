COVERAGEDIR = coverage
ifdef CIRCLE_ARTIFACTS
  COVERAGEDIR = $(CIRCLE_ARTIFACTS)
endif

.DEFAULT_GOAL = all

ifdef VERBOSE
V = -v
else
.SILENT:
endif

.PHONY: all
all: build test cover

.PHONY: install-deps
install-deps:
	glide install

.PHONY: build
build:
	if [ ! -d bin ]; then mkdir bin; fi
	go build $(V) -o bin/gokay

.PHONY: fmt
fmt:
	find . -not -path "./vendor/*" -name '*.go' -type f | sed 's#\(.*\)/.*#\1#' | sort -u | xargs -n1 -I {} bash -c "cd {} && goimports -w *.go && gofmt -w -s -l *.go"

.PHONY: test
test:
	if [ ! -d coverage ]; then mkdir coverage; fi
	go test $(V) ./gkgen -race -cover -coverprofile=$(COVERAGEDIR)/gkgen.coverprofile
	go test $(V) ./gokay -race -cover -coverprofile=$(COVERAGEDIR)/gokay.coverprofile
	go test $(V) ./internal/gkexample -race -cover -coverprofile=$(COVERAGEDIR)/gkexample.coverprofile

.PHONY: cover
cover:
	go tool cover -html=$(COVERAGEDIR)/gkgen.coverprofile -o $(COVERAGEDIR)/gkgen.html
	go tool cover -html=$(COVERAGEDIR)/gokay.coverprofile -o $(COVERAGEDIR)/gokay.html
	go tool cover -html=$(COVERAGEDIR)/gokay.coverprofile -o $(COVERAGEDIR)/gkexample.html

.PHONY: tc
tc: test cover

.PHONY: coveralls
coveralls:
	gover $(COVERAGEDIR) $(COVERAGEDIR)/coveralls.coverprofile
	goveralls -coverprofile=$(COVERAGEDIR)/coveralls.coverprofile -service=circle-ci -repotoken=$(COVERALLS_TOKEN)

.PHONY: clean
clean:
	go clean
	rm -f bin/gokay
	rm -rf coverage/
