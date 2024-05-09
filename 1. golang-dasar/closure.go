package main

import "fmt"


func main() {


  counter := 0

  increment := func() {
    fmt.Println("Hello")
    counter++
  }

  increment()
  increment()
  increment()
  increment()
  increment()
  fmt.Println(counter)
}