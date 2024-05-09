package main

import "fmt"

func main() {

  // person := map[string]string{
  //   "name": "Muhamad Isro Sabanur",
  //   "address": "Penajam",
  // }


  book := make(map[string]string)
  book["title"] = "Habis Gelap Terbitlah Terang"

  // book[]

  // fmt.Println(person)
  // fmt.Println(book)

  if book["title"] != "" {
    fmt.Println(book["title"])
  }

  if length := len(book); length >= 1 {

    fmt.Println("Panjang hanya", length)
  } else {
    fmt.Println("Panjang asd") 
  }
  
  panjang := len(book)

  switch  {
  case panjang >= 1:
    fmt.Println("Length", 1)
  default:
    fmt.Println("Tidak Ada Length")
  }

  // fmt.Println(person["name"])
  // fmt.Println(person["address"])
}