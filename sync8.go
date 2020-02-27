package main

import (
	"fmt"
	"strings"
)

// This version of the word histogram shows the use
// of pipeline pattern to implement the solution.

var data = []string{
	"The yellow fish swims slowly in the water",
	"The brown dog barks loudly after a drink from its water bowl",
	"The dark bird of prey lands on a small tree after hunting for fish",
}

type histogram struct {
	total int
	freq  map[string]int
}
