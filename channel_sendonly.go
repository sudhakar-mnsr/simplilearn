package main

import (
   "fmt"
)

func makeEvenNums(count int, in chan<- int) {
   for i := 1; i <= count; i++ {
      in <- 2 * i
   }
}
func main() {
ch := make(chan int, 10)
makeEvenNums(4,ch)

fmt.Println(<-ch)
fmt.Println(<-ch)
fmt.Println(<-ch)
fmt.Println(<-ch)
}
