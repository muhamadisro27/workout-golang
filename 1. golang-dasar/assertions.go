package main

import "fmt"

type Map map[string]string

func main() {


  result := random()

  resultString := result.(string)

  fmt.Println(resultString)

  // resultInt := result.(int)
  // fmt.Println(resultInt)
  // fmt.Println(result.(type))
  switch value := result.(type) {

  case string:
    fmt.Println("String", value)
   case int:
    fmt.Println("Int", value)
  default:
    fmt.Println("Unknown")
    
  }
  
}

func random() interface{} {
  return "OK"
}