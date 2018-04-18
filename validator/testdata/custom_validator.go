package main

import (
	"fmt"

	"github.com/stratumn/go-indigocore/cs"
)

// Valid validates the transition towards the "init" state
func Valid(l *cs.Link) error {
	fmt.Println("youhou")
	return nil
}

// Invalid validates the transition towards the "init" state
func Invalid(l *cs.Link) error {
	fmt.Println("youhou")
	return nil
}

func main() {}
