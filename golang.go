package main

import "fmt"

func main() {
	/**
	 * ===============================================
	 * In Go, an array is a numbered sequence of
	 * elements of a specific lenght.
	 * ===============================================
	 */

	/**
	 * ===============================================
	 * Here we create an array a that wild hold
	 * exactly 5 ints. The type of elements and
	 * length are both part of the array's type.
	 * By default an array is zero-valued, which for
	 * means 0s.
	 * ===============================================
	 */
	var a [5]int
	fmt.Println("emp:", a)

	fmt.Println("============================================")

	/**
	 * ===============================================
	 * We can, set a value at an index using the
	 * 'array[index] = value' syntax, and get
	 * a value with 'array[index]'
	 * ===============================================
	 */
	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	fmt.Println("============================================")

	/**
	 * ===============================================
	 * the builtin 'len' returns the length of an
	 * array.
	 * ===============================================
	 */
	fmt.Println("len:", len(a))

	fmt.Println("============================================")

	/**
	 * ===============================================
	 * Use this syntax to declare and initialize an
	 * array in one line.
	 * ===============================================
	 */
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	fmt.Println("============================================")

	/**
	 * ===============================================
	 * Array types are one-dimensonal, but you can
	 * compose types to build multi-dimensional
	 * data structures.
	 * ===============================================
	 */
	var arr2D [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			arr2D[i][j] = i + j
		}
	}
	fmt.Println("2d: ", arr2D)
}
