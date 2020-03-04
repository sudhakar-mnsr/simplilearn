package main

import "fmt"

import "net/http"

type msg string

func (m msg) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
   resp.Header().Add("Content-Type", "text/html")
   resp.WriteHeader(http.StatusOK)
   fmt.Fprint(resp, m)
}

func main() {
   msgHandler := msg("Hello from high above!")
   http.ListenAndServe(":4040", msgHandler)
}
