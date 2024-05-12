package main


import (
  "fmt"
)

type validationError struct {
  Message string
}

type notFoundError struct {
  Message string
}

func (v *validationError) Error() string {
  return v.Message
}

func (v *notFoundError) Error() string {
  return v.Message
}

func main() {


  err := SaveData("", 1)

  fmt.Println(err)
  
  
}

func SaveData(id string, data interface{}) error {

  if id == "" {
    return &validationError{Message: "ID cannot be empty"}
  }

  if id != "isro" {
    return &notFoundError{Message: "Data not found"}
  }

  return nil
  
}

