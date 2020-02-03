// Sample program to show see if the class can find the bug.
package main

import (
	"fmt"
	"log"
)

// customError is just an empty struct.
type customError struct{}

// Error implements the error interface.
func (c *customError) Error() string {
	return "Find the bug."
}
