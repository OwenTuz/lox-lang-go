package interpreter

import (
  "bufio"
  "fmt"
  "log"
  "os"
  "strings"
)

func RunFile(filename string) {
  inputFile, err := os.Open(filename)
  if err!= nil {
    log.Fatal(err)
  }
  defer inputFile.Close()

  source := bufio.NewScanner(inputFile)
  run(source)
}

func Prompt() {
  prompt := bufio.NewScanner(os.Stdin)
  scanStdinLoop:
  for {
    fmt.Print("-> ")
    if prompt.Scan() {
      line := bufio.NewScanner(strings.NewReader(prompt.Text()))
      run(line)
    } else {
      switch err := prompt.Err(); err {
      case nil:
        // we received EOF
        fmt.Println()
        break scanStdinLoop
      default:
        log.Fatal(err)
      }
    }
  }
}

func run(source *bufio.Scanner) {
  source.Split(bufio.ScanRunes)
  for source.Scan() {
    fmt.Println(source.Text())
  }
}
