package main

import (
	"fmt"
	"net"
	"os"
	"strings"
	"time"

	curr "github.com/vladimirvivien/learning-go/ch11/curr0"
)

var currencies = curr.Load("./data.csv")

func main() {
   ln, err := net.Listen("tcp", ":4040")
   if err != nil {
      fmt.Println(err)
      os.Exit(1)
   }
   defer ln.Close()
   fmt.Println("Global Currency Service... Listening on port 4040")
   
   // Connection-loop - handle incoming requests
   for {
      conn, err := ln.Accept()
      if err != nil {
         fmt.Println(err)
         conn.Close()
         continue
      }
      fmt.Println("Connected to ", conn.RemoteAddr())
      // delegate connection to a goroutine
      go handleConnection(conn)
   }
}
