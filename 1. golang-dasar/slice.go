package main

import (
  "fmt"
)

func main() {

  // name := [...]string{"Eko", "Kurniawan", "Khannedy", "Joko", "Anwar", "Roozy","Baka","Yaro"}

  // slice := name[:]

  // var slice4  []string = name[4:]

  newSlice := make([]string, 2, 2)

  newSlice[0] = "eko"
  newSlice[1] = "kurniawan"
  new := append(newSlice, "Reza")
  
  fmt.Println(newSlice, new)
  fmt.Println(cap(newSlice), len(newSlice))
  fmt.Println(cap(new), len(new))
  
}