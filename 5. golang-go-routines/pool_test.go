package golang_go_routine

import(
  "testing"
  "fmt"
  "sync"
  "time"
)

func TestPool(t *testing.T) {
  var pool sync.Pool = sync.Pool{
    New: func()interface{} {
      return "new"
    },
  }

  pool.Put("Muhamad")
  pool.Put("Isro")
  pool.Put("Sabanur")


  for i := 0; i<10; i++ {
    go func() {
      data := pool.Get()
      fmt.Println(data)

      time.Sleep(time.Second * 1)
      pool.Put(data)
    }()
  }

  time.Sleep(time.Second * 11)
}
