//Write a simple asynchronous producer-consumer program. It should meet the following requirements:

// There are one producer which would generate a series of unique integers.
// The producer should produce the messages faster than any consumer.
// There are 2 consumers would consume the generated integers (such as print them to stdout).
// Each integer can only be consumed once.
package main

import (
	"fmt"
)

func main() {
	numChan := make(chan int)
	go processMessages(numChan)
	go processMessages(numChan)

	for i := 0; i < 10; i++ {
		numChan <- i
	}
	fmt.Scanln()
}

func processMessages(ch <-chan int) {
	for val := range ch {
		fmt.Println(val)
	}
}
