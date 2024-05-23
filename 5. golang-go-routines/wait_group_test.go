package golang_go_routine

import(
  "testing"
  "fmt"
  "time"
  "sync"
)

func RunAsynchronous(group *sync.WaitGroup) {
  defer group.Done()

  group.Add(1)

  fmt.Println("Hello")
  time.Sleep(time.Second * 1)
  
}

func TestWaitGroup(t *testing.T) {
  group := &sync.WaitGroup{}

  for i:= 0; i<100; i++ {
    go RunAsynchronous(group)
  }

  group.Wait()
  fmt.Println("Complete")
}