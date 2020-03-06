/* Ping
 */
package main

import (
        "bytes"
        "fmt"
        "io"
        "net"
        "os"
)

const myIPAddress = "127.0.0.1"
const ipv4HeaderSize = 20

func main() {
   if len(os.Args) != 2 {
      fmt.Println("Usage: ", os.Args[0], "host")
      os.Exit(1)
   }
   
   localAddr, err := net.ResolveIPAddr("ip4", myIPAddress)
   if err != nil {
      fmt.Println("Resolution error", err.Error())
      os.Exit(1)
   }
   
   remoteAddr, err := net.ResolveIPAddr("ip4", os.Args[1])
   if err != nil {
      fmt.Println("Resolution error", err.Error())
      os.Exit(1)
   }
   
   conn, err := net.DialIP("ip4:icmp", localAddr, remoteAddr)
   checkError(err)
   
   var msg [512]byte
   msg[0] = 8 //echo
   msg[1] = 0 // code 0
   msg[2] = 0 // checksum, fix later
   msg[3] = 0 // checksum, fix later
   msg[4] = 0 // identifier [0]
   msg[5] = 13 // identifier[1] (arbitrary)
   msg[6] = 0 // sequence[0]
   msg[7] = 37 // sequence[1] (arbitrary)
   len := 8
   
   // now fix checksum bytes
   check := checkSum(msg[0:len])
   msg[2] = byte(check >> 8)
   msg[3] = byte(check & 255)
   
   // send the message
   _, err = conn.Write(msg[0:len])
   checkError(err)
   
   fmt.Print("Message sent:   ")
   for n := 0; n < 8; n++ {
      fmt.Print(" ", msg[n])
   }
   fmt.Println()
   
   // receive a reply
   size, err2 := conn.Read(msg[0:])
   checkError(err2)
   
   fmt.Print("Message received:")
   for n := ipv4HeaderSize; n < size; n++ {
      fmt.Print(" ", msg[n])
   }
   fmt.Println()
   os.Exit(0)
}
