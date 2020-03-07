/* SaveJSON
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
