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
 
func New(text string) error {
   return &errorString{text}
}

func main() {
   if err := webCall(); err != nil {
      fmt.Println(err)
      return 
   }
   fmt.Println("Life is good")
}

func webCall() error {
   return New("BAD Request")
}
