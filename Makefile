build:
	go build  -o build ./cmd/todo.go

install:
	go install -v ./cmd/todo.go

.PHONY: build
