GO := godep go
APP := alone

test:
	$(GO) test ./...

build:
	$(GO) build -o $(APP)

run:
	$(GO) run main.go 