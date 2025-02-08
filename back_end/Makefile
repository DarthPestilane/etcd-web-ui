CGO_ENABLED=0
GO111MODULE=on

build:
	go build -ldflags='-s -w' -v -o bin/api .

build-linux:
	GOOS=linux GOARCH=amd64 $(MAKE) build

tidy:
	go mod tidy
