package main

import "fmt"

func sum(numbers []int, ch chan int) {
	total := 0
	for _, num := range numbers {
		total += num
	}
	ch <- total
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	ch := make(chan int)
	go sum(numbers[:len(numbers)/2], ch)
	go sum(numbers[len(numbers)/2:], ch)

	// collect results
	sum1 := <-ch
	sum2 := <-ch

	fmt.Println("total sum: ", sum1+sum2)
}
