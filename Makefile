
build:
	go build -o server cmd/server/main.go

run: generate build
	./server

run/just: build
	./server

generate: 
	go generate ./...

.PHONY: build run
