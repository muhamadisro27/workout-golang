package main

import "fmt"


func main() {
  runApplication()
}

func endApp() {
  fmt.Println("End App !")

  message := recover()

  fmt.Println("Terjadi Error: ", message)
}

func runApplication() {
  defer endApp()

  panic("error while fetching api ###")

  fmt.Println("Run Application")
}