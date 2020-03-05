package main

import (
   "encoding/json"
   "fmt"
   "net/http"
   "curr1"
)

var currencies = curr1.Load("./data.csv")

