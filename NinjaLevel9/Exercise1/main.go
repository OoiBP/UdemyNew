package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("CPU:\t\t", runtime.NumCPU())
	fmt.Println("Go routine:\t\t", runtime.NumGoroutine())
	wg.Add(2)
	go routine1()
	go func() {
		fmt.Println("From routine 2")
		wg.Done()
	}()
	fmt.Println("Go routine:\t\t", runtime.NumGoroutine())
	wg.Wait()
	fmt.Println("Go routine:\t\t", runtime.NumGoroutine())

}

func routine1() {
	fmt.Println("From routine 1")
	wg.Done()
}

/*
in addition to the main goroutine, launch two additional goroutines
	each additional goroutine should print something out
use waitgroups to make sure each goroutine finishes before your program exists
*/
