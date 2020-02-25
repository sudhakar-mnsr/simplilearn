package main

import (
   "fmt"
)
func count(start, stop, delta int) {
   for i := start; i <= stop; i += delta {
      fmt.Println(i)
   }
}

func main() {
   // starts := []int{10,40,70,100}
   
   go func(start, stop, delta int) {
      for i := start; i <= stop; i += delta {
         fmt.Println(i)
      }
   }(100, 150, 10)
   
   go func(start, stop, delta int) {
      for i := start; i <= stop; i += delta {
         fmt.Println(i)
      }
   }(10, 50, 10)

   go count(150,200,10)
   fmt.Scanln()
}
