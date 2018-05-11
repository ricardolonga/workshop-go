package main

import (
	"fmt"
	"time"
)

func hello() {
	fmt.Println("Hello world goroutine.")
}

func _main() {
	go hello()

	fmt.Println("Main function.")

	time.Sleep(1 * time.Second)
}
