package main

import (
   "fmt"
   "sync"
)

const MAX = 1000
const workers = 2

func main() {
values := make(chan int)
result := make(chan int, workers)
var wg sync.WaitGroup

go func() {
   defer close(values)
   for i := 1; i < MAX; i++ {
      if (i%3) || (i%5) {
         values <- i
      }
   } 
}()



