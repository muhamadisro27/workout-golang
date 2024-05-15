package main

import (
  "fmt"
  "reflect"
)

type Person struct {
  Name string `required:"true" max:"10"`
  Email string `required:"true" max:"10"`
  Address string `required:"true" max:"10"`
}

func readField(value any) {
  valueType := reflect.TypeOf(value)
  valueValue := reflect.ValueOf(value)

  for i := 0; i < valueType.NumField(); i++ {
    valueName := valueType.Name()
    field := valueType.Field(i)
    fieldValue := valueValue.Field(i)
    fieldName:= valueType.Field(i).Name

    required := field.Tag.Get("required")
    max := field.Tag.Get("max")
    
    

    fmt.Printf("%s: %v %v %v \n", field.Name, fieldValue.Interface(), fieldName, valueName)

    fmt.Println(required,max)
  }
}

func IsValid(value any) (result bool) {
  result = true
  t := reflect.TypeOf(value)

  for i :=0; i < t.NumField(); i++ {
    f := t.Field(i)
    
    if f.Tag.Get("required") == "true" {
      data := reflect.ValueOf(value).Field(i).Interface()

      result = data != ""

      if result == false {      
        return result
      }
    }

  }
    return result
}

func main() {

  var person Person = Person{
    Name: "Roozy",
    Email: "oqibz@example.com",
    Address: "Test",
  }

  fmt.Println(IsValid(person))

}