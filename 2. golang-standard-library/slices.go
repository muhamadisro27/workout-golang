package main

import (
  "fmt"
  "slices"
)

func main() {

  names := []string{"Roozy", "Qt", "Hi"}
  values := []int{1, 2, 3}


  fmt.Println(slices.Max(values))

}