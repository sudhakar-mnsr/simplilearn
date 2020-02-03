package main

import (
   "fmt"
   "net/temporary"
)

// temporary is declared to test for the existence of the method coming
// from the net package.
type temporary interface {
	Temporary() bool
}

func main() {
   var c *client
   line, err := c.reader.ReadString('\n')
   e := err.(type)
   e.Temporary()
}
