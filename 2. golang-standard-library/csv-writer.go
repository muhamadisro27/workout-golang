package main

import (
  "encoding/csv"
  "os"
)

func main() {

 writer := csv.NewWriter(os.Stdout)

 _ = writer.Write([]string{"Roozy", "Qt", "Az"})
 _ = writer.Write([]string{"Roozy", "Qt", "Az"})

  writer.Flush()
  
}