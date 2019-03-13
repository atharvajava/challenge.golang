package main

import "sync"

var wg sync.WaitGroup //It syncronize the multiple go routines
func main() {
	problem1()

	problem2()
}
