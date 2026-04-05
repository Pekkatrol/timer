package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func check_error() (value int) {
    args := os.Args

	value, err := strconv.Atoi(args[1])
	if err != nil {
		print_error("You must put an integer as parameter\n")
	}
	if value <= 0 {
		print_error("You must put a positive non-zero integer\n")
	}
	return
}

func timer_loop() {
    var value int = check_error()
	remaining := time.Duration(value) * time.Second
	for remaining > 0 {
		fmt.Printf("\r\033[KTimer: %s", remaining)
		time.Sleep(time.Second)
		remaining -= time.Second
	}

	fmt.Printf("\r\033[KTimer: 0s\n")
}
