package main

import "fmt"

type Map map[string]string

func main() {


  isro := NewMap("Reza")


  fmt.Println(isro)
  
}

func NewMap(name string) Map {
  if name == "" {
    return nil
  }

  return Map{
    "name": name,
  }
}
