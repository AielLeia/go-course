package main

import "fmt"

func main() {
	/**
	 * ==================================
	 * The most basic type,
	 * with a single condition
	 * ==================================
	 */
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i++
	}

	fmt.Println("===================================")

	/**
	 * ==================================
	 * A classic inital/condition/after
	 * 'for' loop
	 * ==================================
	 */
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	fmt.Println("===================================")

	/**
	 * ==================================
	 * 'for' without a conditon will loop
	 * repeatedly until you 'break' out
	 * of teh loop or 'return" from
	 * the enclosing funtion
	 * ==================================
	 */
	for {
		fmt.Println("loop")
		break
	}

	fmt.Println("===================================")

	/**
	 * ==================================
	 * You can also 'continue" to the
	 * next iteration of the loop
	 * ==================================
	 */
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
