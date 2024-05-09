package main

import "fmt"

type Map map[string]string

type Address struct {
  City, Province, Country string
}

func main() {

  var address1 Address = Address{"Subang", "Jawa Barat", "Indonesia"}
  var address2 *Address = &address1
  address1.City = "Bandung"

  address2.City = "Malang"

  // address2 = &Address{"Jakarta", "DKI Jakarta", "Indonesia"}
  *address2 = Address{"Jakarta", "DKI Jakarta", "Indonesia"}

  fmt.Println(address1)
  fmt.Println(address2)
  
}
