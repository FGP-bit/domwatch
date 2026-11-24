BINARY  := domwatch
VERSION ?= dev
LDFLAGS := -s -w -X main.version=$(VERSION)

.PHONY: run build test fmt vet clean

## run: ejecuta el CLI contra un dominio (make run DOMAIN=example.com)
DOMAIN ?= example.com
run:
	go run ./cmd/domwatch check $(DOMAIN)

## build: compila el binario en ./bin
build:
	go build -ldflags "$(LDFLAGS)" -o bin/$(BINARY) ./cmd/domwatch

## test: corre los tests
test:
	go test ./...

## fmt / vet: formato y análisis estático
fmt:
	go fmt ./...
vet:
	go vet ./...

## clean: borra los binarios
clean:
	rm -rf bin
