package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"curr1"
)

// currency service client (json)
// sends request as curr1.CurrencyRequest
// receives []curr1Curency
