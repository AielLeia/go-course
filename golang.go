package main

import "fmt"

func main() {
	/**
	 * ==========================================================
	 * Closing a channel incdicates that no more value will be
	 * send ont it. This can be useful to communicate completion
	 * to the channel's receivers.
	 * ==========================================================
	 */

	/**
	 * ==========================================================
	 * In this example we'll use a jobs channelsto communicate
	 * work to be done from the main() goroutine to a worker
	 * goroutine. When we have no more jobs for the worker
	 * we'll close the jobs channels.
	 * ==========================================================
	 */
	jobs := make(chan int, 5)
	done := make(chan bool)

	/**
	 * ==========================================================
	 * Here's the worker goroutine. It repeadtedly receives from
	 * jobs with j, more := <-jobs. In this special 2-value from
	 * of receive, the more value will be false if jobs has been
	 * closed and all value in the channel have already been
	 * received. We use this to notify on done when we've worked
	 * all our jobs.
	 * ==========================================================
	 */
	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	/**
	 * ==========================================================
	 * This sends 3 jobs to the worker over the jobs channel,
	 * then closes it.
	 * ==========================================================
	 */
	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")

	/**
	 * ==========================================================
	 * We awiat the worker using the synchronization approach
	 * we saw earlier.
	 * ==========================================================
	 */
	<-done
}
