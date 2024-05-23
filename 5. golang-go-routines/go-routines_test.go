package golang_go_routine

import(
  "testing"
  "time"
  "fmt"
)

func HelloWorld() {
  fmt.Println("Hello World")
}

func TestHelloWorld(t *testing.T) {

  go HelloWorld()

  fmt.Println("Ups")

  time.Sleep(time.Second * 1)
  
}

func DisplayNumber(number int) {
  fmt.Println("Display number ", number)
}

func TestDisplayNumber(t *testing.T) {
  for i:=1; i <= 10000; i++ {
     go DisplayNumber(i)
  }

  time.Sleep(time.Second *5 )
}