package golang_go_routine

import(
  "testing"
  "fmt"
  "sync"
)

func AddToMap(data *sync.Map, value int, group *sync.WaitGroup) {

  defer group.Done()
  
  group.Add(1)
  
  data.Store(value, value)
}

func TestMap(t *testing.T) {
  data := &sync.Map{}
  group := &sync.WaitGroup{}

  
  for i :=0; i<100; i++ {
    AddToMap(data, i, group)
  }

  group.Wait()
  fmt.Println("Complete")

  data.Range(func(key, value interface{}) bool {
    fmt.Println(key, ":", value)
    return true
  })
}