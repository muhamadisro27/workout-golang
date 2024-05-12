package main

import (
  "fmt"
  "flag"
)

func main() {


  var host *string = flag.String("host", "localhost", "db")
  var username *string = flag.String("username", "root", "username")
  var password *string = flag.String("password", "root", "password")
  var port *int = flag.Int("port", 3306, "custom port")

  flag.Parse()

  fmt.Println(*host, *username, *password, *port)

}