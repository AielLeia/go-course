package main

import "fmt"

func main() {
	/**
	 * =================================================
	 * Slices are a key data type in Go, giving a more
	 * powerful interface to sequences than arrays.
	 * =================================================
	 */

	/**
	 * =================================================
	 * Unlike arrays, slices are typed only by the
	 * elements they contain (not the number of elements).
	 * To create an empty slice witn*h non-zero length,
	 * use the buildin 'make'. Here we make a slice of
	 * 'string's of length 3 (initally zero-valued)
	 * =================================================
	 */
	s := make([]string, 3)
	fmt.Println("emp:", s)

	fmt.Println("====================================================")

	/**
	 * =================================================
	 * We can set and get just like with arrays.
	 * =================================================
	 */
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("====================================================")

	/**
	 * =================================================
	 * 'len' returns the length of the slice as
	 * expected.
	 * =================================================
	 */
	fmt.Println("len:", len(s))

	fmt.Println("====================================================")

	/**
	 * =================================================
	 * In addtion to these basic operations, slices
	 * supprot several more that make the richer than
	 * arrays. One is the buildin 'append", wich returns
	 * a clice containing one or more new values.
	 * Note that we need to accept a return value
	 * from append as we may get a new slice value.
	 * =================================================
	 */
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	fmt.Println("====================================================")

	/**
	 * =================================================
	 * Slices can also be 'copy'd. Here we create an
	 * slice 'c' of the same length as 's' and copy
	 * into 'c' from 's'.
	 * =================================================
	 */
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	fmt.Println("====================================================")

	/**
	 * =================================================
	 * Slices support a 'slice' operator with the syntax
	 * 'slice[low:high]'. For example, this gets a slice
	 * of th elements 's[2]', 's[3]' and 's[4]'
	 * =================================================
	 */
	l := s[2:5]
	fmt.Println("sl1:", l)

	fmt.Println("====================================================")

	/**
	 * =================================================
	 * This clices up to (but excluding) 's[5]'
	 * =================================================
	 */
	l = s[:5]
	fmt.Println("sl2:", l)

	fmt.Println("====================================================")

	/**
	 * =================================================
	 * And this slices up from (and including) 's[2]'
	 * =================================================
	 */
	l = s[2:]
	fmt.Println("sl3:", l)

	fmt.Println("====================================================")

	/**
	 * =================================================
	 * We can declare and initialize a variable
	 * for slice in a single line as well.
	 * =================================================
	 */
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	fmt.Println("====================================================")

	/**
	 * =================================================
	 *
	 * =================================================
	 */
	arr2D := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		arr2D[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			arr2D[i][j] = i + j
		}
	}
	fmt.Println("2d:", arr2D)
}
