package main

import "fmt"

func main() {
	func() {
		fmt.Println("Hello from anonymous!")
	}()
}

/*
Build and use an anonymous func
*/
