GO := godep go
APP := alone
PKG := github.com/ckeyer/alone
VERSION := $(shell cat ./VERSION.txt)
LD_FLAGS := -X $(PKG)/version.version=$(VERSION) -w

test:
	$(GO) test ./...

build:
	$(GO) build -ldflags="$(LD_FLAGS)" -o $(APP)

run:
	$(GO) run main.go 
