package main

import (
	"os"
	"fmt"
)

func print_error(text string) {
    fmt.Fprint(os.Stderr, text)
    os.Exit(2)
}