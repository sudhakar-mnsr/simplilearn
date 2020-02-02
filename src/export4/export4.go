package main

import (
   "fmt"
   "export4/users"
)

func main() {
   u := users.Manager{Title: "Director"}
   u.Name = "sudhakar"
   u.ID = 10
   fmt.Println(u)
}
