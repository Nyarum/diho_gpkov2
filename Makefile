
build:
	go build -o server cmd/server/main.go

run: build
	./server

.PHONY: build run