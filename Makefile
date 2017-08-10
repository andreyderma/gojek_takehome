SHELL := /bin/bash

# The name of the executable (default is current directory name)
TARGET := $(shell echo $${PWD\#\#*/})
.DEFAULT_GOAL: $(TARGET)

# These will be provided to the target
VERSION := 1.0.0
BUILD := `git rev-parse HEAD`

# go source files, ignore vendor directory
SRC = $(shell find . -type f -name '*.go' -not -path "./vendor/*")

.PHONY: all clean build uninstall fmt simplify check run

all: uninstall check build

$(TARGET): $(SRC)
	@go build -o $(TARGET)

clean:
	@rm -f $(TARGET)

build:
	@go build -o $(TARGET)

uninstall: clean
	@rm -f $$(which ${TARGET})

fmt:
	@go fmt -w $(SRC)

simplify:
	@go fmt -s -w $(SRC)

check:
	@test -z $(shell go fmt main.go | tee /dev/stderr) || echo "[WARN] Fix formatting issues with 'make fmt'"
	@for d in $$(go list ./... | grep -v /vendor/); do go fmt $${d}; done
	@go test ./... --cover -v

run: build
	@$(TARGET)