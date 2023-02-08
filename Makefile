VERSION := $(shell git rev-parse --short HEAD)
BUILDTIME := $(shell date -u '+%Y-%m-%dT%H:%M:%SZ')

GOLDFLAGS += -X main.Version=$(VERSION)
GOLDFLAGS += -X main.Buildtime=$(BUILDTIME)
GOFLAGS = -ldflags "$(GOLDFLAGS)"

.PHONY: build
build:
	@go build -o ./bin/api $(GOFLAGS) ./cmd/web

.PHONY: run
run:
	@go run ./cmd/web

.PHONY: clean
clean:
	@rm -rf ./bin