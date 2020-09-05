package main

import "fmt"

/**
 * ====================================================================
 * The '(int, int)' in this function signature showns that the call
 * fucntion returns 2 'int's
 * ====================================================================
 */
func vals() (int, int) {
	return 3, 7
}

func main() {
	/**
	 * ====================================================================
	 * here we use the 2 different return value from the call with
	 * multiple assignment.
	 * ====================================================================
	 */
	a, b := vals()
	fmt.Println(a, b)

	fmt.Println("================================================================")

	/**
	 * ====================================================================
	 * If you only want a subset of the returned values, use the blank
	 * identifier _.
	 * ====================================================================
	 */
	_, c := vals()
	fmt.Println(c)
}
