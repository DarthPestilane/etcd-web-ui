CGO_ENABLED=0

build:
	go build -ldflags='-s -w' -v -o bin/api .

build-linux:
	GOOS=linux $(MAKE) build

run-yecao: build
	./api start --endpoint=103.143.81.64:12379 --username=eva --password=mhXVPazGg2T77Ae7zvIDp5BruSak19Heq --key-prefix=/eva/

tidy:
	go mod tidy
