VERSION_MAJOR  ?= 0
VERSION_MINOR  ?= 1
VERSION_BUILD  ?= 999
VERSION_TSTAMP ?= $(shell date -u +%Y%m%d-%H%M%S)
VERSION_SHA    ?= $(shell git rev-parse --short HEAD)
VERSION ?= v$(VERSION_MAJOR).$(VERSION_MINOR).$(VERSION_BUILD)-$(VERSION_TSTAMP)-$(VERSION_SHA)

GOLANGCILINT_VER ?= 1.42.1
GOLANGCILINT_BINARY ?= $(GOPATH)/bin/golangci-lint
GOLANGCILINT_CHECK = $(if $(filter $(shell go env GOHOSTOS),darwin),,golangci)

MODPATH ?= $(shell head -n1 go.mod | grep module | awk '{print $$2}')
BINNAME ?= negra

BUILD_FLAGS_DD ?= $(if $(filter $(shell go env GOHOSTOS),darwin),,-d)
BUILD_FLAGS ?= -ldflags="$(BUILD_FLAGS_DD) -s -w -X $(MODPATH)/pkg/config.version=$(VERSION)" -tags netgo -installsuffix netgo


$(shell mkdir -p ./out)

.PHONY: env-info
env-info:
	@echo
	id
	@echo
	@echo PWD=$(shell pwd)
	@env | grep -i GO
	@echo
	@go version

.PHONY: install-tools
install-tools:
	apk add openssh-client git coreutils

.PHONY: build
build: out/$(BINNAME) post-build-check

.PHONY: out/$(BINNAME)
out/$(BINNAME): env-info
	CGO_ENABLED=1 go build $(BUILD_FLAGS) -a -o $@ cmd/negra.go

.PHONY: post-build-check
post-build-check:
	ls -l out/
	@test -f out/$(BINNAME)

.PHONY: version
version:
	@echo $(VERSION)

.PHONY: clean
clean:
	rm -rf out/

.PHONY: golangci
golangci:
	@golangci-lint --version 2>&1 > /dev/null || echo "...............will be (re-)installed" && \
	  wget -O-  -q https://github.com/golangci/golangci-lint/releases/download/v$(GOLANGCILINT_VER)/golangci-lint-$(GOLANGCILINT_VER)-linux-amd64.tar.gz | \
	  tar xzf - -O golangci-lint-$(GOLANGCILINT_VER)-linux-amd64/golangci-lint > $(GOLANGCILINT_BINARY) && \
	  chmod 755 $(GOLANGCILINT_BINARY)

.PHONY: test
test: env-info $(GOLANGCILINT_CHECK)
	@gofmt -d  $(shell find ./pkg ./cmd -name '*.go')
	CGO_ENABLED=1 go vet $(BUILD_FLAGS) ./...
	CGO_ENABLED=1 golangci-lint run
	gocritic version && CGO_ENABLED=1  gocritic check ./... || true
	CGO_ENABLED=1 go test $(BUILD_FLAGS) ./...

.PHONY: run-local
run-local: env-info
	CGO_ENABLED=1 go run $(BUILD_FLAGS) cmd/negra.go -v 4
	# --config ./config.yaml --zap-time-encoding=iso8601
