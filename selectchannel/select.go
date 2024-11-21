package main

import (
	"fmt"
	"time"
)

func ping(ch chan string) {
	time.Sleep(2 * time.Second)
	ch <- "Ping"
}

func pong(ch chan string) {
	time.Sleep(1 * time.Second)
	ch <- "Pong"
}

func main() {
	pi := make(chan string)
	po := make(chan string)

	go ping(pi)
	go pong(po)

	for i := 0; i < 2; i++ {
		select {
		case msg := <-pi:
			fmt.Println("Received", msg)
		case msg := <-po:
			fmt.Println("received", msg)
		}
	}
}
