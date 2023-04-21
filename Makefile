all: clean build test test-race lint gofumpt docker-image docker-run
APP_NAME := backrunner

GOPATH := $(if $(GOPATH),$(GOPATH),~/go)
VERSION := $(shell git describe --tags --always)

clean:
	rm -rf ${APP_NAME} build/

build:
	go build -trimpath -ldflags "-X main._BuildVersion=${VERSION}" -v -o ${APP_NAME} cmd/main.go

test:
	go test ./...

test-race:
	go test -race ./...

lint:
	gofmt -d -s .
	gofumpt -d -extra .
	go vet ./...
	staticcheck ./...
	golangci-lint run

gofumpt:
	gofumpt -l -w -extra .

docker-image:
	DOCKER_BUILDKIT=1 docker build --platform linux/amd64 --progress=plain  --build-arg VERSION=${VERSION} . -t ${APP_NAME}-${VERSION}:${VERSION}

osx-docker-image:
	DOCKER_BUILDKIT=1 docker build --platform linux/arm64  --progress=plain  --build-arg APP_NAME=${APP_NAME} --build-arg VERSION=${VERSION} . -t ${APP_NAME}-${VERSION}:latest

docker-run:
	docker run --env-file=.env.example backrunner-${VERSION}
