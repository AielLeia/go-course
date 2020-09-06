package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	/**
	 * ===================================================
	 * Suppose we have a function call f(s). Here's how
	 * we'd call that in the usual way, running it
	 * synchronously.
	 * ===================================================
	 */
	f("direct")

	/**
	 * ===================================================
	 * To invoke this function in a goroutine, use go f(s)
	 * This new goroutine will execute concurrently with
	 * the calling on.
	 * ===================================================
	 */
	go f("goroutine")

	/**
	 * ===================================================
	 * You can also start a goroutine for an anonymous
	 * function call.
	 * ===================================================
	 */
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	/**
	 * ===================================================
	 * When we run this program, we see the output of
	 * the blocking call first, the the interleaved output
	 * of the tw goroutines. This interleaving reflect
	 * teh routines being run concurrently by Go runtinme.
	 * ===================================================
	 */
	time.Sleep(time.Second)
	fmt.Println("Done")
}
