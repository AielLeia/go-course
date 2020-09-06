package main

import "fmt"

func main() {
	/**
	 * ===========================================================
	 * By default channels are unbuffered, meaning that the will
	 * only accept sends (chan <-) if there is a corresponding
	 * receive (chan <-) ready to receive the sen value. Buffurred
	 * channels accept a limited number of value without a
	 * corresponding receiver for those values.
	 * ===========================================================
	 */

	/**
	 * ===========================================================
	 * Here we make a channel of strings buffering up to 2 values.
	 * ===========================================================
	 */
	messages := make(chan string, 2)

	/**
	 * ===========================================================
	 * Because this channel is buffered, we can send these values
	 * into the channel without a corresponding concurrent receive.
	 * ===========================================================
	 */
	messages <- "buffered"
	messages <- "channel"

	/**
	 * ===========================================================
	 * Later we can receive the two values as usual.
	 * ===========================================================
	 */
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
