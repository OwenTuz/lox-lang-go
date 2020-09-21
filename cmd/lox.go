package main

import (
	"github.com/owentuz/lox-lang-go/pkg/interpreter"
	"log"
	"os"
)

func main() {
	if len(os.Args) > 2 {
		log.Fatalf("Usage: lox [script]")
	}
	if len(os.Args) == 2 {
		interpreter.RunFile(os.Args[1])
	} else {
		interpreter.Prompt()
	}
}
