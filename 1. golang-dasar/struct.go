package main

import "fmt"


type Customer struct  {
  Name, Address string
  Age int
}

func main() {


  var customer Customer

  customer.Name = "Isro"
  customer.Address = "Jakarta"
  customer.Age = 23

  customer2 := Customer{
    Name : "Roozy",
    Address : "Bandung",
    Age : 23,
  }

  fmt.Println(customer)
  fmt.Println(customer2)
}
