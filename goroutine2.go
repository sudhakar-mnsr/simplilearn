package main

import (
   "fmt"
)
func count_star(start, stop, delta int) {
   for i := start; i <= stop; i += delta {
      fmt.Println("***",i,"***")
   }
}
func count(start, stop, delta int) {
   for i := start; i <= stop; i += delta {
      fmt.Println(i)
   }
}

func main() {
   starts := []int{10,40,70,100}
   
   // Since j is updated in each iteration it is impossible to determine
   // the value read for j by closure. 
   for _, j := range starts {
      go count(j, j+20, 10)
   }
   // Since computation of closure putting us in non deterministic value
   // it is safe to pass the value as parameter for goroutine
   for _, j := range starts {
      go func(s int) {
         count_star(s, s+20, 10)
      }(j)
   }
   
   fmt.Scanln()
}
