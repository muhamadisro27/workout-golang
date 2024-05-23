package golang_go_routine

import(
  "testing"
  "time"
  "fmt"
  "sync"
  "strconv"
)

func TestChannel(t *testing.T) {

  channel := make(chan string)
  
  defer close(channel)

  go func() {
    time.Sleep(2 * time.Second)
    channel <- "Roozya"

    fmt.Println("Success sent to channel")
  }()

  var data string
  
  data = <- channel

  fmt.Println(data)

}

func TestChannelParams(t *testing.T) {


  channel := make(chan string)
  channel2 := make(chan string)

  defer close(channel)
  defer close(channel2)

  go GiveMeResponse(channel) 
  go GiveMeResponse2(channel2) 

  data := <- channel
  data2 := <- channel2

  fmt.Println(data)
  fmt.Println(data2)

}



func GiveMeResponse2(channel chan string) {


  time.Sleep(5 * time.Second)
  channel <- "Roozyqt2"

}

func GiveMeResponse(channel chan string) {


  time.Sleep(2 * time.Second)
  channel <- "Roozyqt"

}

func OnlyIn(channel chan<- string) {


  time.Sleep(2 * time.Second)
  channel <- "Roozyqt"

}

func OnlyOut(channel <-chan string) {


  data := <- channel

  fmt.Println(data)

}

func TestOnlyInOut(t *testing.T) {

  channel := make(chan string)
  defer close(channel)

  go OnlyIn(channel)
  go OnlyOut(channel)

  time.Sleep(time.Second * 5)  
}

func TestBuffered(t *testing.T) {
  channel := make(chan string, 5)

  defer close(channel)

  go func (){
    for i:=1; i<=5; i++ {
      channel <- fmt.Sprintf("Roozy %d", i)
    }
  }()

  go func() {
    fmt.Println(<-channel)
    fmt.Println(<-channel)
    fmt.Println(<-channel)
    fmt.Println(<-channel)
    fmt.Println(<-channel)
  }()

  time.Sleep(time.Second * 2)
  fmt.Println("Selesai")
}

func TestRange(t *testing.T) {

  channel := make(chan string)

  go func() {
    for i:=1; i<=10; i++ {
        channel <- "Perulangan ke-" + strconv.Itoa(i)
    }
    close(channel)
  }()


  for data := range channel {
    fmt.Println(data)
  }

  fmt.Println("selesai")
  
}

func TestPipeline(t *testing.T) {
  ch1 := make(chan int)
  ch2 := make(chan string)

  go func(){
    for i:= 1; i < 10; i++ {
      ch1 <- i
    }
    close(ch1)
  }()

  go func() {
    for v:= range ch1 {
      number := v *2 

      timeNow := time.Now()

      data := fmt.Sprintf("%d diterima pada waktu %d:%d:%d", number, timeNow.Hour(), timeNow.Minute(), timeNow.Second())
      
      ch2 <- data
    }
    close(ch2)
  }()

  for result := range ch2 {
    fmt.Println(result)
  }
}

func TestSelectChannel(t *testing.T) {
  channel1 := make(chan string)
  channel2 := make(chan string)
  defer close(channel1)
  defer close(channel2)

  go GiveMeResponse(channel1)
  go GiveMeResponse(channel2)

  counter := 0;
  
  for {
    select {
      case data:= <- channel1:
        fmt.Println("data dari channel 1", data)
        counter++
      case data:= <- channel2:
        fmt.Println("data dari channel 2", data)
        counter++
      default:
      fmt.Println("Menunggu data...")
    }

    if counter == 2 {
      break
    }
  }
}

func TestRaceCondition(t *testing.T) {

  x:= 0

  for i:=1; i<= 1000; i++ {
    go func(){
      for j := 1; j <=100; j++{
        x = x+1
      }
    }()
  }

  time.Sleep(time.Second * 5)
  if x != 100000 {
    t.Error("x tidak sama dengan 100000")
  }
  fmt.Println("Counter : ", x)
  
}

func TestRaceConditionMutex(t *testing.T) {

  x:= 0
  var mutex sync.Mutex

  for i:=1; i<= 1000; i++ {
    go func(){
      for j := 1; j <=100; j++{
        mutex.Lock()
        x = x+1
        mutex.Unlock()
      }
    }()
  }

  time.Sleep(time.Second * 5)
  if x != 100000 {
    t.Error("x tidak sama dengan 100000")
  }
  fmt.Println("Counter : ", x)

}