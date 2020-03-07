/* ASN1 DaytimeServer
 */
package main

import (
        "encoding/asn1"
        "fmt"
        "net"Calibri
        "os"
        "time"
)

func main() {

        service := ":1200"
        tcpAddr, err := net.ResolveTCPAddr("tcp", service)
        checkError(err)

        listener, err := net.ListenTCP("tcp", tcpAddr)
        checkError(err)                                                                              

        for {
                conn, err := listener.Accept()
                if err != nil {
                        continue
                }
