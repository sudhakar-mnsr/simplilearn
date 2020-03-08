/* LoadJSON
 */

package main

import (
        "encoding/json"
        "fmt"
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
