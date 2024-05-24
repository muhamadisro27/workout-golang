package golang_context

import (
  "testing"
  "context"
  "fmt"
)

func TestContext(t *testing.T) {

  background := context.Background()
  fmt.Println("background :",background)

  todo := context.TODO()
  fmt.Println("TODO",todo)
  
}