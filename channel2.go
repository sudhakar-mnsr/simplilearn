package main

import (
   "fmt"
)

func main() {
   ch := make(chan int, 15)
   
   ch <- 1
   ch <- 2
   ch <- 3
   ch <- 4
   ch <- 5
   ch <- 6
   ch <- 7
   ch <- 8
   ch <- 9
   ch <- 10
   ch <- 11
   ch <- 12
   ch <- 13
   close(ch)
   
   for i := 1; i <= 15; i++ {
      if val, opened := <-ch; opened {
         fmt.Println(val)
      } else {
         fmt.Println("Channel closed")
      }
   }
} 
