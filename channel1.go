package main

import (
   "fmt"
)

func main() {
   var ch chan int
   // Without the below line it is deadlock with errors
   ch = make(chan int)
   
   go func() {
      ch <- 12
   }()
   fmt.Println(<-ch)
}
