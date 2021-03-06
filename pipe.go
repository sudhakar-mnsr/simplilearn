package main

import (
   "fmt"
   "strings"
)

var data = []string{
	"The yellow fish swims slowly in the water",
	"The brown dog barks loudly after a drink from its water bowl",
	"The dark bird of prey lands on a small tree after hunting for fish",
}

type histogram struct {
   total int
   freq map[string]int
}

func (h histogram) ingest() <-chan string {
   out := make(chan string)
   go func() {
      defer close(out)
      for _, line := range data {
         out <- line
      }
   }()
   return out
}

func (h histogram) split(in <-chan string) <-chan string {
   out := make(chan string)
   go func() {
      defer close(out)
      for line := range in {
         words := strings.Split(line, " ")
         for _, word := range words {
            out <- strings.ToLower(word)
         }
      }
   }()
   return out
}

func (h histogram) count(in <-chan string) chan struct{} {
   done := make(chan struct{})
   go func() {
      defer close(done) 
      for word := range in {
         h.freq[word]++
         h.total++
      }
   }()
   return done
}

func main() {
   h := histogram{freq: make(map[string]int)}
   done := make(chan struct{})
   go func() {
      defer close(done)
      <-h.count(h.split(h.ingest()))
   }()
   <-done
   fmt.Println(h.freq)
}  
