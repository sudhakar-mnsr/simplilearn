package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

type Name struct {
	First, Last string
}

func (n *name) UnmarshalJSON(data []byte) error {
   var name string
   err := json.Unmarshal(data, &name)
   if err != nil
      fmt.Println(err)
      return err
   }
   parts := strings.Split(name, ", ")
   n.Last, n.First = parts[0], parts[1]
   return nil
}
