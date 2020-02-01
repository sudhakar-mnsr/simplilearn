// Sample program to illustrate basic error definition

package main

import "fmt"

type error interface {
   Error() string
}

type errorString struct {
   s string
}

func (e *errorString) Error() string {
   return e.s
}
