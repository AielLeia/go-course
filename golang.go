package main

import "fmt"

func main() {
	/**
	 * ======================================================
	 * Create a new channel with make(chan val-type). Channels
	 * are typed by the values the convey.
	 * ======================================================
	 */
	messages := make(chan string)

	/**
	 * ======================================================
	 * Send a value into channel using the channel <- syntax.
	 * Here we send ping to the message channel we made
	 * above, from a new goroutine.
	 * ======================================================
	 */
	go func() {
		messages <- "ping"
	}()

	/**
	 * ======================================================
	 * The <- channel syntyax receives a value from the
	 * channel. Here we'll receive the ping message we sent
	 * above and print it out.
	 * ======================================================
	 */
	msg := <-messages
	fmt.Println(msg)

	/**
	 * ======================================================
	 * When we run the program the ping message is
	 * successfully passed from one goroutine to another via
	 * our channel.
	 *
	 * By default sends and reveives block until both the
	 * sender and receiver are ready. This property allowed
	 * us to wait at the end of our program for the ping
	 * message without having to use any other syncronization.
	 * ======================================================
	 */
}
