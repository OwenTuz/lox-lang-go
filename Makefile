.PHONY:  build clean run test

build:
	go build -o bin/lox cmd/lox.go

clean:
	rm -rf ./bin/*

run:
	go run cmd/lox.go $(ARGS)

test:
	go test $(ARGS) ./... 
