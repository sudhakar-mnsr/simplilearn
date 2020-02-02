package main

import (
   "fmt"
   "export1/counters"
)

func main() {
   counter := counters.AlertCounter(10)
   fmt.Printf("Counter: %d\n", counter)
}
