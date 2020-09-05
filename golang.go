package main

import "fmt"

func zeroval(ival int) {
	/**
	 * ====================================================================
	 * We'll show how pointers work in contrast to values with 2 function:
	 * zeroval and zeroptr. zeroval has an int parameter, so arguments will
	 * be passed to it by value. zeroval will get a copy of ival distinct
	 * from the on in the calling function.
	 * ====================================================================
	 */
	ival = 0
}

func zeroptr(iptr *int) {
	/**
	 * ====================================================================
	 * zeroptr in contrast has an *int parameter, meaning that it takes an
	 * int pointer. The *iptr code in the function body then dereferences
	 * the pointer from its memory address to current value at that address.
	 * Assigning a valur a dereferenced pointer changes the value at the
	 * referenced address.
	 * ====================================================================
	 */
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	/**
	 * ====================================================================
	 * the &i syntax gives the memory address of i, i.e. a pointer
	 * to i.
	 * ====================================================================
	 */
	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	/**
	 * ====================================================================
	 * Pointer can be printed too.
	 * ====================================================================
	 */
	fmt.Println("pointer:", &i)
}
