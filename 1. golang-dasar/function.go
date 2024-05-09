package main

import "fmt"

type BlackList func(string) bool

func main() {


  luas, keliling := persegi_panjang(4,6)

  fmt.Println("Luas persegi panjang adalah ...", luas)
  fmt.Println("Keliling persegi panjang adalah ...", keliling)


  sum := sumAll(10,20,30,40,50,60,10)

  numbers := []int{30,20,40,10,50,60}


  fmt.Println(sumAll(numbers...))
  fmt.Println(sum)

  good_bye := goodBye

  fmt.Println(good_bye("Roozy"))

  sayHelloWithFilter("Anjing", spamFilter)

  // fmt.Println(sayHello)
  filter := spamFilter

  sayHelloWithFilter("Isro", filter)

  blackList := func(name string) bool {
    if name == "Anjing" {
      return true
    } else {
      return false
    }
  }
  
  registerUser("Anjinga", blackList)
  
}

func registerUser(name string, blacklist BlackList) {

  if blacklist(name) {
    fmt.Println("Your account is blocked !" , name)
  } else {
    fmt.Println("Welcome", name)
  }
}

func goodBye(name string) string {

  return "Good Bye " + name
}

type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
  fmt.Println("Hello ", filter(name))
}

func spamFilter(name string) string {
  if name == "Anjing" {
    return "..."
  } else {
    return name
  }
}

func sumAll(numbers ...int) int {

  total := 0

  for _, number := range numbers {
    total += number
  }

  return total
  
}

func persegi_panjang(panjang int, lebar int) (luas int, keliling int) {

  luas = panjang * lebar

  keliling = 2 * (panjang + lebar)


  return luas, keliling
  
}