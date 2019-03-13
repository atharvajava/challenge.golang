package main

import (
	"log"
	"math/rand"
	"time"
)

//var mux sync.Mutex   ----- In this scenario i would prefer using mutex over channels as channels have a performance penalty as compared to mutex
// Plus I dont think we actually need all the overhead of  amazing features of channels for this problem
// But for the problem i have written both mutex and channels either can be used here
// channels are good since it provides inbuilt thread safety and internally they also use mutex

func problem2() {
	log.Printf("problem2: started --------------------------------------------")

	//
	// Todo:
	//
	// Throttle all go subroutines in a way,
	// that every one second one random number
	// is printed.
	//

	for inx := 0; inx < 10; inx++ {
		wg.Add(1)
		done := make(chan bool, 1) // Buffered channel of type boolean and size 1
		go printRandom2(done, inx)
		<-done // Waiting to receive notification from channel
	}

	//
	// Todo:
	//
	// Remove this quick and dirty sleep
	// against a synchronized wait until all
	// go routines are finished.
	//
	// Same as problem1...
	//

	//time.Sleep(5 * time.Second)
	wg.Wait() // Since its already syncronized in this case i dont think this is need still just for the problem i am using this here
	log.Printf("problem2: finished -------------------------------------------")
}

func printRandom2(done chan bool, slot int) {
	defer wg.Done()
	defer func() {
		done <- true // Notifying channel then
	}()
	for inx := 0; inx < 10; inx++ {
		//mux.Lock()
		log.Printf("problem2: slot=%03d count=%05d rand=%f", slot, inx, rand.Float32())
		time.Sleep(1 * time.Second)
		//mux.Unlock()

	}
}
