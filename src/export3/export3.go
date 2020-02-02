package main

import (
   "fmt"
   "export3/users"
)

func main() {
   u := users.User{
      Name: "Chole",
      ID: 10,
   }
   fmt.Printf("User: %#v\n", u)
}
