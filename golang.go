package main

import "fmt"

/**
 * =============================================================
 * This ping functin only accepts a channel for sending values.
 * It would be a compile time error to try to receive on this
 * channel.
 * =============================================================
 */
func ping(pings chan<- string, msg string) {
	pings <- msg
}

/**
 * =============================================================
 * the pong function accepts one channel for receives (pings)
 * and a second for sends (pongs).
 * =============================================================
 */
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	/**
	 * =============================================================
	 * When using channels as funcion parameters, you can specify if
	 * a channel is meant to only send or receive values. This
	 * specificity increases the type safety of the program.
	 * =============================================================
	 */

	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
