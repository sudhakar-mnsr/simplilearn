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

   for {
      fmt.Print(prompt, "> ")
      _, err := fmt.Scanf("%s %s", &cmd, &param)
      if err != nil {
         fmt.Println("Usage: GET <search string or *>")
         continue
      }
      cmdLine := fmt.Sprintf("%s %s", cmd, param)
      if n, err := conn.Write([]byte(cmdLine)); n == 0 || err != nil {
         fmt.Println(err)
         return
      }
      conn.SetReadDeadline(time.Now().Add(time.Millisecond * 5000))
      conbuf := bufio.NewReaderSize(conn, 1024)
      for {
         str, err := conbuf.ReadString('\n')
         if err != nil {
            break
         }
         fmt.Print(str)
         conn.SetReadDeadline(time.Now().Add(time.Millisecond * 700))
      }
   }
}
