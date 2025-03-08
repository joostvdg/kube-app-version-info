
LOCAL_VERSION = $(shell git describe --tags --always)
PACKAGE_VERSION ?= "0.1.0-$(LOCAL_VERSION)"
NAME := kube-app-version-info
MAIN_GO :=

.PHONY: build
build:
	CGO_ENABLED=0 go build -o bin/$(NAME) cmd/server/main.go

.PHONY: buildc
buildc:
	CGO_ENABLED=0 go build -o bin/$(NAME) cmd/client/main.go
