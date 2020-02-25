package main

import (
	"fmt"
	"strings"
)

func main() {
   data := []string{
      "The yellow fish swims slowly in the water",
      "The brown dog barks loudly after a drink from its water bowl",
      "The dark bird of prey lands on a small tree after hunting for fish",
   }
   
   histogram := make(map[string]int)
   wordch := make(chan string)
   
   go func() {
      defer close(wordch)
      for _, line := range data {
         words := strings.Split(line, " ")
         for _, word := range words {    
            word = strings.ToLower(word)
            wordch <- word
         }
      }
   }()
   
   // process the word stream and count words
   // loop until wordsch is closed
   for {
      word, opened := <-wordch
      if !opened {
         break
      }
      histogram[word]++
   }
   
   for k, v := range histogram {
      fmt.Printf("%s\t(%d)\n", k, v)
   }
}
