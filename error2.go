// Sample program to show how to use error variables to help the 
// caller determine the exact error being returned.
package main

import (
   "fmt"
   "errors"
)

var (
   ErrBadRequest = errors.New("Bad Request")
   ErrPageMoved = errors.New("Page Moved")
)

func main() {
   if err := webCall(); err != nil {
      switch err {
      case ErrBadRequest:
         fmt.Println("Bad Request error occurred")
         return
      case ErrPageMoved:
         fmt.Println("Page Moved error") 
         return
      default:
         fmt.Println(err)
         return
      }
   }
   fmt.Println("Life good")
}
