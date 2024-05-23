package golang_go_routine

import(
  "testing"
  "fmt"
  "sync"
  "sync/atomic"
  // "time"
)


func TestAtomic(t *testing.T) {

  var x int64 = 0
  group := sync.WaitGroup{}

  for i:=1; i<= 1000; i++ {
    go func(){
      group.Add(1)
      for j := 1; j <=100; j++{
        atomic.AddInt64(&x, 1)
      }
      group.Done()
    }()
  }

  group.Wait()
  if x != 100000 {
    t.Error("x tidak sama dengan 100000")
  }
  fmt.Println("Counter : ", x)

}