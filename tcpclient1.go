package main

import (
   "bufio"
   "fmt"
   "net"
   "time"
)

var host, port = "127.0.0.1", "4040"
var addr = net.JoinHostPort(host, port)
var deadline = time.Now().Add(time.Millisecond * 700)

const prompt = "curr"
const buffLen = 1024

func main() {
   dailer := &net.Dailer{
      Timeout: time.Second * 300,
      KeepAlive: time.Minute * 5,
   }
   conn, err := dailer.Dial("tcp", addr)
   if err != nil {
      fmt.Println(err)
      return
   }
   defer conn.Close()
   fmt.Println("Connected to Global Currency Service")
   var cmd, param string

