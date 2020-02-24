package main

import (
   "encoding/json"
   "fmt"
   "os"
   "time"
)

type Name struct {
   First, Last string
}

func (n *Name) MarshalJSON() ([]byte, error) {
   return []byte(fmt.Sprintf("\"%s, %s\"", n.Last, n.First)), nil
}

type Book struct {
	Title       string
	PageCount   int
	ISBN        string
	Authors     []Name
	Publisher   string
	PublishDate time.Time
}
