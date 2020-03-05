package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/vladimirvivien/learning-go/ch11/curr1"
)

// Simple json server for currency service
// As in tcpserver example, returns []curr1.Currency
// This time, however, data is encoded as JSON
// Test with:
// curl-X POST -d '{"get":"Euro"}' http://localhost:4040/currency

var currencies = curr1.Load("./data.csv")
