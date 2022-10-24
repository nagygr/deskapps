.phony: all chk

chk:
	go fmt ./cmd/...
	go vet ./cmd/...

all: chk
	go build -o build/deskapps ./cmd/deskapps
