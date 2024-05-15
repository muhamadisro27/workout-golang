package main

import (
  "fmt"
  "bufio"
  "strings"
  "io"
)

func main() {


  input := strings.NewReader("this is long string\nwith new line\n")

  reader := bufio.NewReader(input)

  for {
    line, _, err := reader.ReadLine()

    if err == io.EOF {
      break
    }

    fmt.Println(string(line))
  }
}