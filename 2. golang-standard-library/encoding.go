package main

import (
  "fmt"
  "encoding/base64"
  "encoding/csv"
  "io"
  "strings"
)

func main() {

  var encoded = base64.StdEncoding.EncodeToString([]byte("Muhamad Isro Sabanur"))

  fmt.Println(encoded)

  var decoded, err = base64.StdEncoding.DecodeString(encoded)

  if err != nil {
    fmt.Println(err.Error())
  } else {
    fmt.Println(string(decoded))
  }

  csvString := "Muhamad,Isro,Sabanur\n" + "Roozy,Roo,Qt\n"

  reader := csv.NewReader(strings.NewReader(csvString))

  for {
    record, err := reader.Read()

    if err == io.EOF {
      break
    }

    fmt.Println(record)
  }

}