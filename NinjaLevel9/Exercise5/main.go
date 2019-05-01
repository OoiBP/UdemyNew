package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var counter int32 = 0
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
	atomic.AddInt32(&counter, 1)
	wg.Done()
}

/*
Fix the race condition you created in exercise #4 by using package atomic
*/
