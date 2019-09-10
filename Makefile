VERSION := 0.1.0
SERVICE_NAME := $(shell grep "^module" go.mod | rev | cut -d "/" -f1 | rev)

export GO111MODULE=on

.PHONY: install
install: build

.PHONY: build
build:
	CGO_ENABLED=0 go build -o bin/notifyer \
        -ldflags "-X main.version=$(VERSION) -X main.serviceName=$(SERVICE_NAME)" \
        ./cmd/notify

.PHONY: container
container:
	@docker build . \
		-t notifyer:$(VERSION) \
		--build-arg VERSION=$(VERSION) \
		--build-arg SERVICE_NAME=$(SERVICE_NAME) \
