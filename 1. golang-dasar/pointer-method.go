package main

import "fmt"

type Map map[string]string

type Man struct {
  Name string
}

func (man *Man) Married() {
  man.Name = "Mr. " + man.Name
}

func main() {

  eko := Man{"Isro"}
  eko.Married()

  fmt.Println(eko)
}