package main

import (
   "encoding/gob"
   "fmt"
   "os"
   "time"
)

type Name struct {
   First, Last string
}

type Book struct {
   Title string
   PageCount int
   ISBN string
   Authors []Name
   Publisher string
   PublishDate time.Time
}

func main() {
books := []Book{
   Book{
      Title: "Learning Go",
      PageCount: 375,
      ISBN: "123"
      Authors: []Name{{"Vladimir", "Vivien"},{"MNSR", "Sudhakar"},},
      Publisher: "Packt",
      PublishDate: time.Date(2016, time.July, 0, 0, 0, 0, 0, time.UTC),
   },
   Book{
      Title: "The Go Programing Language",
      PageCount: 380,
      ISBN: "456"
      Authors: []Name{{"Brian", "Kernighan"},{"MNSR", "Sudhakar"},{"Alan", "Donavan"}},
      Publisher: "Addison-Wesley",
      PublishDate: time.Date(2015, time.October, 26, 0, 0, 0, 0, time.UTC),
   },
   Book{
      Title: "Introducing Go",
      PageCount: 124,
      ISBN: "8910"
      Authors: []Name{{"Caleb", "Doxsey"},{"MNSR", "Sudhakar"}},
      Publisher: "OReilly",
      PublishDate: time.Date(2016, time.January, 26, 0, 0, 0, 0, time.UTC),
   },
