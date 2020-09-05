package main

import "fmt"

/**
 * =============================================================
 * Here's a function that takes two 'int's and returns their
 * sum as an int.
 * =============================================================
 */
func plus(a int, b int) int {
	/**
	 * =============================================================
	 * Go requires explicit returns, i.e it won't automactilly
	 * return the value of the last expression.
	 * =============================================================
	 */
	return a + b
}

/**
 * =============================================================
 * When you have multiple consecutive parameters of the same
 * type, you may omit the type name for the like-typed
 * parameters up to the final parameter that declares the types.
 * =============================================================
 */
func plusPlus(a, b, c int) int {
	return a + b + c
}

func main() {
	/**
	 * =============================================================
	 * Call a function juste as you'd expect, with 'name(args)'
	 * =============================================================
	 */
	res := plus(1, 2)
	fmt.Println("1 + 2 = ", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1 + 2 + 3 = ", res)
}
