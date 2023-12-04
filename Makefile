.PHONY: default server client deps fmt clean all release-all contributors

BUILDTAGS=debug
SERVER_RELEASE=bin/ngroxd
CLIENT_RELEASE=bin/ngrox

default: all
	
test:
	go test -race --coverprofile=coverage.coverprofile --covermode=atomic ./...

deps: 
	go mod download

server: deps
	go build -tags '$(BUILDTAGS)' -o $(SERVER_RELEASE) cmd/ngroxd/ngroxd.go

fmt:
	go fmt internal/...

client: deps
client: 
	go build -tags '$(BUILDTAGS)' -o $(CLIENT_RELEASE) cmd/ngrox/ngrox.go

release-client: BUILDTAGS=release
release-client: CLIENT_RELEASE=$(ARG_CLIENT_RELEASE)
release-client: client

release-server: BUILDTAGS=release
release-client: SERVER_RELEASE=$(ARG_SERVER_RELEASE)
release-server: server

release-all: fmt release-client release-server

all: fmt client server

contributors:
	echo "Contributors to ngrox, both large and small:\n" > CONTRIBUTORS
	git log --raw | grep "^Author: " | sort | uniq | cut -d ' ' -f2- | sed 's/^/- /' | cut -d '<' -f1 >> CONTRIBUTORS
