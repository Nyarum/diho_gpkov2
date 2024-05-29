
build:
	go build -o server cmd/server/main.go

run: generate build
	./server

generate: 
	go generate ./...

.PHONY: build run
