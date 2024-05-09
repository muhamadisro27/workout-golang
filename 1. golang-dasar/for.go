package main

import "fmt"

func main() {

  // counter := 1

  // for counter <= 10 {
  //   fmt.Println("Cetak angka-", counter)
  //   counter++
  // }

  for counter:=1; counter <= 10; counter++ {
    fmt.Println("Cetak angka-", counter)
  }

  names := []string{"Muhamad", "Isro", "Sabanur"}

  for i := 0; i < len(names); i++ {
    fmt.Println(names[i])
  }


  for _, value := range names{
    // fmt.Println("key", key, "=", value)
    fmt.Println(value)
  }


  for counter:=1; counter <= 10; counter++ {
    if counter % 2 != 0 {
      continue
    }

    fmt.Println("Perulangan ke-", counter)
  }
}