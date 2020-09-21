.PHONY:  build clean run test gofmt

build:
	go build -o bin/lox cmd/lox.go

clean:
	rm -rf ./bin/*

run:
	go run cmd/lox.go $(ARGS)

test:
	go test $(ARGS) ./... 

gofmt:
	gofmt -w $(shell find . -type f -name '*.go')
