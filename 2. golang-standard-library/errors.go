package main

import (
  "fmt"
  "errors"
)

var(
  ValidationError = errors.New("Validation Error")
  NotFoundError = errors.New("Not Found Error")
)

func main() {

  err := GetById("a")

  if err != nil {
     if errors.Is(err, ValidationError) {
       fmt.Println("validation error!")
     } else if errors.Is(err, NotFoundError) {
       fmt.Println("not found error!")
     } else {
       fmt.Println("unknown error!")
     }
  }

}

func GetById(id string) error {

  if id == "" {
    return ValidationError
  }

  if id != "isro" {
    return NotFoundError
  }

  return nil

}