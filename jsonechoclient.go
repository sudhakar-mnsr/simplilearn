/* JSON EchoClient
 */
package main

import (
        "bytes"
        "encoding/json"
        "fmt"
        "io"
        "net"
        "os"
)

type Person struct {
        Name  Name
        Email []Email
}

type Name struct {
        Family   string
        Personal string
}

type Email struct {
        Kind    string
        Address string                                                
}
