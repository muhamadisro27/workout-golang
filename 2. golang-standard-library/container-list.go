package main

import (
  "fmt"
  "container/list"
)

func main() {

  var data *list.List = list.New()

  data.PushBack("rozzy")
  data.PushBack("joko")
  data.PushBack("eko")

  // fmt.Println(data.Front().Value)
  // fmt.Println(data.Front().Next().Value)
  // fmt.Println(data.Back().Value)

  for e := data.Front(); e != nil; e = e.Next() {
    fmt.Println(e.Value)
  }

}