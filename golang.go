package main

import (
	"fmt"
	"time"
)

/**
 * =============================================================
 * This is the function we'll run in a goroutine. The done
 * channel will be used to notify another goroutine that this
 * function's work is done.
 * =============================================================
 */
func worker(done chan bool) {

	fmt.Print("working...")
	time.Sleep(time.Second * 2)
	fmt.Println("done")

	/**
	 * =============================================================
	 * Send a valuen to notify the we're done.
	 * =============================================================
	 */
	done <- true
}

func main() {
	/**
	 * =============================================================
	 * We can use channels to synchronize execution across goroutine
	 * Here's an example of using a blocking receive to wait for
	 * goroutine to finish. When waiting for multiple goroutine
	 * to finish, yous may prefer to use a WaitGroup.
	 * =============================================================
	 */

	/**
	 * =============================================================
	 * Start a worker goroutine, giving it the channel to notify on.
	 * =============================================================
	 */
	var done chan bool = make(chan bool, 1)
	go worker(done)

	/**
	 * =============================================================
	 * Block until we receice a notification from the worker on the
	 * channel.
	 *
	 * If you removed the <- done line from this program, the
	 * program would exit before the worker even started.
	 * =============================================================
	 */
	<-done
}
