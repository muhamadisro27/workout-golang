package main

import (
  "os"
  "io"
  "bufio"
  // "fmt"
)

func readFile(name string)(string,error) {
  file, err := os.OpenFile(name, os.O_RDONLY, 0666)
  if err != nil {
    return "",err
  }

  defer file.Close()

  reader := bufio.NewReader(file)

  var message string

  for {
    line, _, err := reader.ReadLine()
    if err == io.EOF {
      break
    }

    message += string(line) + "\n"
    
  }

  return message, nil
}

func createNewFile(name, message string) error {
  file, err := os.OpenFile(name, os.O_CREATE|os.O_WRONLY, 0666)
  if err != nil {
    return err
  }
  defer file.Close()

  file.WriteString(message)

  return nil

}

func addToFile(name, message string) error {
  file, err := os.OpenFile(name, os.O_RDWR| os.O_APPEND, 0666)
  if err != nil {
    return err
  }

  defer file.Close()

  file.WriteString(message + "\n")
  
  return nil
}

func main() {
  // createNewFile("hello.txt", "Hello World")

  // read, _ := readFile("hello.txt")

 
  // fmt.Println(read)

  addToFile("hello.txt", "Hello World Lagi 2")
 

}