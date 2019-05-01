package main

import (
	"fmt"
	"sync"
)

var counter = 0
var wg sync.WaitGroup

func main() {
	wg.Add(50)
	for i := 0; i < 50; i++ {
		go increment()
	}
	wg.Wait()
	fmt.Println("Counter:", counter)
}

func increment() {
	v := counter
	v++
	counter = v
	wg.Done()
}

/*
Using goroutines, create an incrementer program
	have a variable to hold the incrementer value
	launch a bunch of goroutines
		each goroutine should
			read the incrementer value
				store it in a new variable
			yield the processor with runtime.Gosched()
			increment the new variable
			write the value in the new variable back to the incrementer variable
use waitgroups to wait for all of your goroutines to finish
the above will create a race condition.
Prove that it is a race condition by using the -race flag
*/
