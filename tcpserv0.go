package main

import (
   "fmt"
   "net"
   "os"
   "strings"
   "time"
   curr "curr0"
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
      go handleConnection(conn)
   }
}

func handleConnection(conn net.Conn) {
   defer conn.Close()
   for {
      cmdLine := make([]byte, (1024 * 4))
      n, err := conn.Read(cmdLine)
      if n == 0 || err != nil {
         return
      }

      cmd, param := parseCommand(string(cmdLIne[0:n]))
      if cmd == "" {
         continue
      }

      switch strings.ToUpper(cmd) {
         case "GET":
            result := curr.Find(currencies, param)
            if len(result) == 0 {
               conn.Write([]byte("Nothing found\n"))
               continue
            }
            for _, cur := range result {
               _, err := fmt.Fprintf(conn, "%s %s %s %s\n", curr.Name, curr.Code, curr.Number, cur.Country,)
               if err != nil {
                  return
               }
               conn.SetWriteDeadline(time.Now().Add(time.Second * 5))
            }
            conn.SetWriteDeadline(time.Time{})
         default:
            conn.Write([]byte("Invalid command\n"))
       }
   }
} 

func parseCommand(cmdLine string) (cmd, param string) {
   parts := strings.Split(cmdLine, " ")
   if len(parts) != 2 {
      return "", ""
   }
   cmd = strings.TrimSpace(parts[0])
   param = strings.TrimSpace(parts[1])
   return
}
