package main

import (
	"fmt"
	"time"
)

func main() {
	/**
	 * =============================================================
	 * Timeouts are important for programs that connect to external
	 * resources or that otherwise need to bound execution time.
	 * Implementing timeout in Go is easy elegant thanks to channel
	 * and select.
	 * =============================================================
	 */

	/**
	 * =============================================================
	 * For our example, suppose we're executiong an external call
	 * that returns its result on a channel c1 after 2s. Note that
	 * the channel is buffered, so the send in goroutine is
	 * non-blocking. This a common pattern to prevent goroutine
	 * leaks in case the channel is never read.
	 * =============================================================
	 */
	var c1 chan string = make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()

	/**
	 * =============================================================
	 * Here's the select implementing a timeout. res := <-c1 awaits
	 * the result and <-time.After awaits a value to be send after
	 * teh timeout of 1s. Since select proceeds with the first
	 * receive that's ready, we'll take timeout case if the operation
	 * takes more than the allowed 1s.
	 * =============================================================
	 */
	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1")
	}

	/**
	 * =============================================================
	 * If we allow a longer timeout of 3s, then the receive from c2
	 * will succeed and we'll print the result.
	 * =============================================================
	 */
	var c2 chan string = make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()

	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}
}
