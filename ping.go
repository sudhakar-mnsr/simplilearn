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
