package main

import "fmt"

func main() {
	c := make(chan int)

	for index := 0; index < 10; index++ {
		go func() {
			for jndex := 0; jndex < 10; jndex++ {
				c <- jndex
			}
		}()
	}

	for kndex := 0; kndex < 100; kndex++ {
		fmt.Println(kndex, <-c)
	}
	fmt.Println("About to exit")
}

/*
write a program that
	launches 10 goroutines
		each goroutine adds 10 numbers to a channel
	pull the numbers off the channel and print them
*/
