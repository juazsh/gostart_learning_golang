package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func generateData(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- rand.Intn(100)
	}
	close(ch)
}

func processData(id int, input chan int, output chan int) {
	for data := range input {
		fmt.Printf("Worker %d processing data: %d\n", id, data)
		time.Sleep(500 * time.Millisecond)
		output <- data * 2
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	input := make(chan int)
	output := make(chan int)
	var wg sync.WaitGroup
	go generateData(input)
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			processData(id, input, output)
		}(i)
	}

	go func() {
		wg.Wait()
		close(output)
	}()

	for result := range output {
		fmt.Println("Result:", result)
	}
}
