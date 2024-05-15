package main

import (
  "fmt"
  "time"
)

func main() {

  var duration1 time.Duration = 100 * time.Second
  var duration2 time.Duration = 10 * time.Second

  fmt.Printf("duration1 : %d\n", duration1)
  fmt.Printf("duration2 : %d\n", duration2)

}