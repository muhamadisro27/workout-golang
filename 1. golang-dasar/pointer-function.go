package main

import "fmt"

type Map map[string]string

type Address struct {
  City, Province, Country string
}


func main() {

  var address1 Address = Address{"Subang", "Jawa Barat", ""}

  changeToIndonesia(&address1)

  var name *string
  fill_name := "roozy"

  // age := new()

  name = &fill_name
  
  fmt.Println(*name)

  fmt.Println(address1)
}

func changeToIndonesia(address *Address) {
  address.Country = "Indonesia"
}