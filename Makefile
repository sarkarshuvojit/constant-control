BRANCH := $(shell git rev-parse --abbrev-ref HEAD)
HASH := $(shell git rev-parse --short HEAD)

default:
	@echo "Cmds: [build | run]"

test:
	@go test ./...

build:
	@go build -o .bin/cosntant-control-$(BRANCH)-$(HASH)

build-image:
	@docker build -t sarkarshuvojit/constant-control:local .
