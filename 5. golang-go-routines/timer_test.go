package golang_go_routine

import(
  "testing"
  "fmt"
  "sync"
  "time"
)

func TestTimer(t *testing.T) {
  timer := time.NewTimer(5 * time.Second)
  fmt.Println(time.Now())

  time := <- timer.C

  fmt.Println(time)
}

func TestAfter(t *testing.T) {
  channel := time.After(5 * time.Second)

  tick := <-channel

  fmt.Println(tick)
}

func TestAfterFunc(t *testing.T) {
  group:= sync.WaitGroup{}
  group.Add(1)

  time.AfterFunc(1* time.Second, func() {
    fmt.Println("Execute after 1 second")
    group.Done()
  })
  
  group.Wait()
}