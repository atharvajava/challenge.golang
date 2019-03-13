package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync/atomic"
)

var counter int64 // Just to track if there are only 100 prints

func problem1() {
	ch := make(chan int, 100)
	// initializing a buffered channel of size 100
	// To execute only 100 go routines

	log.Printf("problem1: started --------------------------------------------")

	//
	// Todo:
	//
	// Quit all go routines after
	// a total of exactly 100 random
	// numbers have been printed.
	//
	// Do not change the 25 in loop!
	//

	//This could also be changed to 4 make it generate 100 random numbers
	for inx := 0; inx < 10; inx++ {
		wg.Add(1) //Tells to syncronize the waitgroup here
		go printRandom1(ch, inx)

	}

	for i := 1; i <= 100; i++ {
		ch <- i // Storing values inside the channel
	}
	close(ch) // Closing channel to notify that the job has completed and thus no more values should be sent to it

	//
	// Todo:
	//
	// Remove this quick and dirty sleep
	// against a synchronized wait until all
	// go routines are finished.
	//

	//time.Sleep(5 * time.Second)
	wg.Wait()
	// Tells the compiler to wait till all the go routines are finished
	// This is much better then the hacky time.sleep where we are hardcoding and hoping the program to close on its own
	fmt.Println(counter)
	log.Printf("problem1: finised --------------------------------------------")
}

func printRandom1(ch chan int, slot int) {

	//
	// Do not change 25 into 10!
	//

	defer wg.Done()
	// Using defer here to basically ensure that wg.Done() happens so that if there in error
	// in this function we can tell the waitgroups that we are done and not waiting anymore. like decrementing what we added

	for inx := 0; inx < 25; inx++ {
		_, ok := <-ch // checks if channel has been closed
		if !ok {      // if closed return and end go routine
			return
		}
		log.Printf("problem1: slot=%03d count=%05d rand=%f", slot, inx, rand.Float32())
		atomic.AddInt64(&counter, 1) // the atomic package helps to avoid data race , Mutex can also be used here.
	}
}

// The above should also be achievable using context package
// By adding a check on counter
