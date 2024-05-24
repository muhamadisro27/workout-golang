package golang_context

import (
  "testing"
  "context"
  "fmt"
  "runtime"
  "time"
)

func TestContextValue(t *testing.T) {

  contextA := context.Background()

  contextB := context.WithValue(contextA, "b", "B")
  // contextC := context.WithValue(contextA, "c", "C")

  contextD := context.WithValue(contextB, "d", "D")
  // contextE := context.WithValue(contextB, "e", "E")
  
  // contextF := context.WithValue(contextC, "f", "F")

  contextG := context.WithValue(contextD, "g", "G")

  // fmt.Println(contextA)
  // fmt.Println(contextB)
  // fmt.Println(contextC)
  // fmt.Println(contextD)
  // fmt.Println(contextE)
  // fmt.Println(contextF)
  // fmt.Println(contextG)

  fmt.Println(contextG.Value("b"))
}

func CreateCounter(ctx context.Context) (chan int) {
  destination := make(chan int)

  go func() {
    defer close(destination)
    counter := 1;
    for {
      select {
        case <- ctx.Done():
          return  
        default:
          destination <- counter
          counter++
          time.Sleep(1 * time.Second)
      }
    }
  }()

  return destination
}

func TestContextWithCancel(t *testing.T) {

  fmt.Println("Total Goroutine",runtime.NumGoroutine())
  ctx, cancel := context.WithCancel(context.Background())

  destination := CreateCounter(ctx)

  for n := range destination {
    fmt.Println("Counter",n)
    if n == 10 {
      break
    }
  }
  cancel()

  time.Sleep(2 * time.Second)
  
  fmt.Println("Total Goroutine",runtime.NumGoroutine())
}

func TestContextWithTimeout(t *testing.T) {
  fmt.Println("Total Goroutine",runtime.NumGoroutine())
  ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
  defer cancel()

  destination := CreateCounter(ctx)

  for n := range destination {
    fmt.Println("Counter",n)
  }

  fmt.Println("Total Goroutine",runtime.NumGoroutine())
}

func TestContextWithDeadline(t *testing.T) {
  fmt.Println("Total Goroutine",runtime.NumGoroutine())
  ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(5*time.Second))
  defer cancel()

  destination := CreateCounter(ctx)

  for n := range destination {
    fmt.Println("Counter",n)
  }

  fmt.Println("Total Goroutine",runtime.NumGoroutine())
}