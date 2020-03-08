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

func (p Person) String() string {
        s := p.Name.Personal + " " + p.Name.Family
        for _, v := range p.Email {
                s += "\n" + v.Kind + ": " + v.Address
        }
        return s
}

func main() {
        person := Person{
                Name: Name{Family: "Newmarch", Personal: "Jan"},
                Email: []Email{Email{Kind: "home", Address: "jan@newmarch.name"},
                        Email{Kind: "work", Address: "j.newmarch@boxhill.edu.au"}}}

        if len(os.Args) != 2 {
                fmt.Println("Usage: ", os.Args[0], "host:port")
                os.Exit(1)
        }
        service := os.Args[1]

        conn, err := net.Dial("tcp", service)
        checkError(err)

        encoder := json.NewEncoder(conn)
        decoder := json.NewDecoder(conn)
