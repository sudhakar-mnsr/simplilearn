package main

import (
   "fmt"
   "export2/counters"
)

func main() {
   counter := counters.New(10)
   fmt.Printf("counter value %d \n", counter)
   fmt.Printf("counter type %T \n", counter)
}
