package main

import (
	"errors"
	"fmt"
)

/**
 * ==============================================================
 * By convention, errors are the last return value and have
 * type error, a buit-in interface.
 * ==============================================================
 */
func f1(arg int) (int, error) {
	if arg == 42 {
		/**
		 * ==============================================================
		 * errors.New constructs a basic error value with given error
		 * message.
		 * ==============================================================
		 */
		return -1, errors.New("can't work with 42")
	}

	/**
	 * ==============================================================
	 * A nil value int the error position indicate that there was no
	 * error.
	 * ==============================================================
	 */
	return arg + 3, nil
}

/**
 * ==============================================================
 * It's possioble to use costom types as errors by implementing
 * the Error() method on theme. Here's a variante on the example
 * above that use a custom type to explicitly represent an
 * argument error.
 * ==============================================================
 */
type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		/**
		 * ==============================================================
		 * In this case we use &argError syntax to build a new struct,
		 * supplying values for the two fields arg and prob.
		 * ==============================================================
		 */
		return -1, &argError{arg, "Can't work with it"}
	}
	return arg + 3, nil
}

func main() {
	/**
	 * ==============================================================
	 * The two loo^below test our each of our error-returning
	 * functions. Note that use of an inline errot check on the if
	 * line is a common idiom in Go code.
	 * ==============================================================
	 */
	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}
	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}

	fmt.Println("=========================================================")

	/**
	 * ==============================================================
	 * If you want to programmatically use the data in a custom error
	 * you'll need to get the errot as an instance of the custom
	 * error type via type assertion.
	 * ==============================================================
	 */
	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}
