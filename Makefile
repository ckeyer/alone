GO := go
APP := alone
PKG := github.com/ckeyer/alone
VERSION := $(shell cat ./VERSION.txt)
LD_FLAGS := -X $(PKG)/version.version=$(VERSION) -w

default: local run

test:
	$(GO) test -ldflags="$(LD_FLAGS)" $$(go list ./... |grep -v vendor)

local:
	$(GO) build -ldflags="$(LD_FLAGS)" -o bin/$(APP)

run:
	bin/$(APP)

build386:
	GOOS="linux" \
	GOARCH="386" \
	CGO_ENABLED="0" \
	$(GO) build -ldflags="$(LD_FLAGS)" -o bin/$(APP)