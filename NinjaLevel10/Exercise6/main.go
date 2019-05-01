package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		for index := 0; index < 100; index++ {
			c <- index
		}
		fmt.Println("Closing c...")
		close(c)
	}()

	for v := range c {
		fmt.Println(v)
	}
	fmt.Println("About to exit")
}

/*
write a program that
	puts 100 numbers to a channel
	pull the numbers off the channel and print them
*/
