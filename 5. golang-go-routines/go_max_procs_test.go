package golang_go_routine

import(
  "testing"
  "fmt"
  "runtime"
  "time"
  "sync"
)

func TestGetgomaxprocs(t *testing.T) {

  group:= sync.WaitGroup{}

  for i:= 0; i<100; i++ {
    group.Add(1)
    go func() {
      time.Sleep(3 * time.Second)
    }()
  } 
  
  totalCpu := runtime.NumCPU()

  fmt.Println(totalCpu)

  totalThread := runtime.GOMAXPROCS(-1)
  
  fmt.Println(totalThread)

  totalGoroutine := runtime.NumGoroutine()

  fmt.Println(totalGoroutine)
}