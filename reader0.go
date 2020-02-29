package main

import (
   "fmt"
   "io"
   "os"
)

type alphareader string

func (a alphareader) Read(p []byte) (int, error) {
   count := 0
   for i := 0; i < len(a); i++ {
      if (a[i] >= 'A' && a[i] <= 'Z') ||
         (a[i] >= 'a' && a[i] <= 'z') {
         p[i] = a[i]
      }
      count++
   }
   return count, io.EOF
}

func main() {
   str := alphaReader("Hello! where is the sun?")
   io.Copy(os.Stdout, &str)
   fmt.Println()
}    

