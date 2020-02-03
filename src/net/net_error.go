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
