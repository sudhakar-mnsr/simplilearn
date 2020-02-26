package main

import (
   "fmt"
   "strings"
)

func main() {
   data := []string{ 
      "The yellow fish swims slowly in the water", 
      "The brown dog barks loudly after a drink ...", 
      "The dark bird bird of prey lands on a small ...", 
   }
   wordch := make(chan string)
   histogram := make(map[string]int)

   go func() {
      defer close(wordch)
      for _, line := range data {
         words := strings.Split(line, " ")
         for _, word := range words {
            word = strings.ToLower(word)
            wordch<- word
         }
      }
   }()

   for word := range wordch {
      histogram[word]++
   }
   for k, v := range histogram {
      fmt.Printf("%s\t(%d)\n", k, v)
   }
}
