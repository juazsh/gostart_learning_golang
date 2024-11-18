package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	data := []byte("Ola, Go!")
	err := ioutil.WriteFile("testFile.txt", data, 0644)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println("File created successfully")
}
