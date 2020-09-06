package main

import "fmt"

func main() {
	/**
	 * ===============================================================
	 * In a previous example we saw how for and range provide
	 * iteration over basic data structures. We ca, also use this
	 * syntax to iterate over values recieved from a channel.
	 * ===============================================================
	 */

	/**
	 * ===============================================================
	 * We'll iterate over 2 values in the queue channel.
	 * ===============================================================
	 */
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	/**
	 * ===============================================================
	 * This range iterates over each element as it's received from
	 * queue. Because we closed the channel abose, the iteration
	 * terminates after receiving the 2 elements.
	 * ===============================================================
	 */
	for elem := range queue {
		fmt.Println(elem)
	}
}
