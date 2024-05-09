package main

import "fmt"


type Customer struct  {
  Name, Address string
  Age int
}

type HasName interface {
  GetName() string
}

func main() {

  person := Customer{Name: "Roozy"}

  SayHello(person)
  
}

func (customer Customer) GetName() string {

  return customer.Name
}

func SayHello(hasName HasName) {
  fmt.Println("Hello", hasName.GetName())
}
