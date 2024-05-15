package main

import (
  "fmt"
  "sort"
)

type User struct {
  Name string
  Age int
}

type UserSlice []User

func (userSlice UserSlice) Len() int {
  return len(userSlice)
}

func (userSlice UserSlice) Less(i,j int) bool {
  return userSlice[i].Age < userSlice[j].Age
}

func (userSlice UserSlice) Swap(i,j int) {
  userSlice[i], userSlice[j] = userSlice[j], userSlice[i]
}

func main() {

  var users UserSlice = UserSlice{
    User{Name: "rozzy", Age: 20},
    User{Name: "joko", Age: 30},
    User{Name: "eko", Age: 40},
  }

  sort.Sort(users)

  fmt.Println(users)
}