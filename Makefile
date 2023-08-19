BRANCH := $(shell git rev-parse --abbrev-ref HEAD)
HASH := $(shell git rev-parse --short HEAD)

default:
	@echo "Cmds: [build | run]"

test:
	@go test ./...

build:
	@go build -o .bin/constant-control-$(BRANCH)

run:
	@go build -o .bin/constant-control-$(BRANCH) && ./.bin/constant-control-$(BRANCH)

build-image:
	@docker build -t sarkarshuvojit/constant-control:local .

run-help:
	@go build -o .bin/constant-control-$(BRANCH) && ./.bin/constant-control-$(BRANCH) --help
