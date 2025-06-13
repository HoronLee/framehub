ROOT_DIR    = $(shell pwd)
NAMESPACE   = "default"
DEPLOY_NAME = "template-single"
DOCKER_NAME = "template-single"

include ./hack/hack-cli.mk
include ./hack/hack.mk

GOHOSTOS:=$(shell go env GOHOSTOS)
GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags --always)

ifeq ($(GOHOSTOS), windows)
	#the `find.exe` is different from `find` in bash/shell.
	#to see https://docs.microsoft.com/en-us/windows-server/administration/windows-commands/find.
	#changed to use git-bash.exe to run find cli or other cli friendly, caused of every developer has a Git.
	#Git_Bash= $(subst cmd\,bin\bash.exe,$(dir $(shell where git)))
	Git_Bash=$(subst \,/,$(subst cmd\,bin\bash.exe,$(dir $(shell where git))))
	INTERNAL_PROTO_FILES=$(shell $(Git_Bash) -c "find internal -name *.proto")
	API_PROTO_FILES=$(shell $(Git_Bash) -c "find api -name *.proto")
else
	INTERNAL_PROTO_FILES=$(shell find internal -name *.proto)
	API_PROTO_FILES=$(shell find api -name *.proto)
endif

.PHONY: ctrl
ctrl:
	gf gen ctrl

.PHONY: dao
dao:
	gf gen dao

.PHONY: genall
# Generate gf code
genall:
	gf gen ctrl
	gf gen dao

.PHONY: build
# build
build:
	mkdir -p bin/ && go build -o ./bin/ ./...

.PHONY: run
# run
run:
	gf run main.go