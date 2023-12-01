.PHONY: default server client deps fmt clean all release-all contributors

BUILDTAGS=debug
default: all
	
test:
	go test -race --coverprofile=coverage.coverprofile --covermode=atomic ./...

deps: 
	go mod tidy

server: deps
	go install -tags '$(BUILDTAGS)' ngrok/main/ngrokd

fmt:
	go fmt ngrok/...

client: deps
	go install -tags '$(BUILDTAGS)' ngrok/main/ngrok

release-client: BUILDTAGS=release
release-client: client

release-server: BUILDTAGS=release
release-server: server

release-all: fmt release-client release-server

all: fmt client server

contributors:
	echo "Contributors to ngrok, both large and small:\n" > CONTRIBUTORS
	git log --raw | grep "^Author: " | sort | uniq | cut -d ' ' -f2- | sed 's/^/- /' | cut -d '<' -f1 >> CONTRIBUTORS
