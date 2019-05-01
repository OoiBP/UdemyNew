package main

import (
	"fmt"
	"sync"
)

var counter = 0
var wg sync.WaitGroup
var mutex sync.Mutex

func main() {
	wg.Add(50)
	for i := 0; i < 50; i++ {
		go increment()
	}
	wg.Wait()
	fmt.Println("Counter:", counter)
}

func increment() {
	mutex.Lock()
	v := counter
	v++
	counter = v
	mutex.Unlock()
	wg.Done()
}

/*
Fix the race condition you created in the previous exercise by using a mutex
	it makes sense to remove runtime.Gosched()
*/
