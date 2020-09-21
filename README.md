lox-lang-go
===========

A Golang implementation of the `lox` language from the book [Crafting Interpreters](https://craftinginterpreters.com/) by Bob Nystrom


The original version is written in Java (and reimplemented in C). This is a personal effort to improve both my knowledge of Go and compiler/language-hacking skills in a way that doesn't involve copying code directly from a book.

If you're looking for a book on writing interpreters in Go specifically, try Thorsten Ball's [Writing An Interpreter In Go](https://interpreterbook.com/#get-a-taste).

### Other notes
- The book (and associated source code) uses the term `scanner` but I have chosen `lexer`, mainly to avoid confusion with Golang's [bufio.Scanner](https://golang.org/pkg/bufio/#Scanner)
- I've named Lox errors `Exception`s, probably for the same reason the book's Java implementation calls them `error`: it doesn't clash with the language's built-in name
