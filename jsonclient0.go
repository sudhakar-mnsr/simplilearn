package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"curr1"
)

// currency service client (json)
// sends request as curr1.CurrencyRequest
// receives []curr1Curency

func main() {
var param string
fmt.Print("Currency> ")
_, err := fmt.Scanf("%s", &param)

buf := new(bytes.Buffer)
currRequest := &curr1.CurrencyRequest{Get: param}
err = json.NewEncoder(buf).Encode(currRequest)
if err != nil {
   fmt.Println(err)
   return
}

client := &http.Client{}
req, err := http.NewRequest("POST", "http://127.0.0.1:4040/currency", buf)
if err != nil {
   fmt.Println(err)
   return
}
resp, err := client.Do(req)
if err != nil {
   fmt.Println(err)
   return
}
defer resp.Body.Close()

