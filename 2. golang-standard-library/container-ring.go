package main

import (
  "fmt"
  "strconv"
  "container/ring"
)

func main() {

  var data *ring.Ring = ring.New(5)

  for i := 0; i < data.Len(); i++ {
    data.Value = "Value " + strconv.FormatInt(int64(i), 10)

    data = data.Next()
  }

  data.Do(func(value interface{}) {
    fmt.Println(value)
  })
}