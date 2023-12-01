.PHONY: default server client deps fmt clean all release-all contributors
export GOPATH:=$(shell pwd)

BUILDTAGS=debug
default: all

deps: 
	cd src/ngrok && go mod tidy

server: deps
	cd src/ngrok && go install -tags '$(BUILDTAGS)' ngrok/main/ngrokd

fmt:
	cd src/ngrok && go fmt ngrok/...

client: deps
	cd src/ngrok && go install -tags '$(BUILDTAGS)' ngrok/main/ngrok

release-client: BUILDTAGS=release
release-client: client

release-server: BUILDTAGS=release
release-server: server

release-all: fmt release-client release-server

all: fmt client server

contributors:
	echo "Contributors to ngrok, both large and small:\n" > CONTRIBUTORS
	git log --raw | grep "^Author: " | sort | uniq | cut -d ' ' -f2- | sed 's/^/- /' | cut -d '<' -f1 >> CONTRIBUTORS
