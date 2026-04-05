package main

import (
	"os"
)

func main() {
	args := os.Args

	if (len(args) != 2) {
		print_error("You need one argument\n")
	}
    timer_loop()
}