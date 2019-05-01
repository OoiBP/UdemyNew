package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {
		c <- 42
	}()

	fmt.Println(<-c)
}

// func main() {
// 	c := make(chan int, 1) // Buffered channel

// 	c <- 42

// 	fmt.Println(<-c)
// }
