package main

import (
	"fmt"
	"time"
)

func printMessage(msg string) {
	for i := 1; i <= 5; i++ {
		fmt.Println(msg, i)
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	go printMessage("Running Go Routine")
	printMessage("This is from main")
}
