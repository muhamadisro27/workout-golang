package main

import (
  "fmt"
  "time"
)

func main() {

  fmt.Println(time.Now())

  utc := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)

  fmt.Println(utc.Local())

  parse, _ := time.Parse(time.RFC3339, "2006-01-02T15:04:05Z")
  
  fmt.Println(parse)
}