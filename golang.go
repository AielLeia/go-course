package main

import (
	"fmt"
	"time"
)

func main() {
	/**
	 * ==================================================
	 * Here's a basic switch
	 * ==================================================
	 */
	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	fmt.Println("===================================")

	/**
	 * ==================================================
	 * You can use commas to separate multiple
	 * expression in the same case statement.
	 * We use the optional 'default' case in this example
	 * as well.
	 * ==================================================
	 */
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	fmt.Println("===================================")

	/**
	 * ==================================================
	 * 'switch' without an expression is an alternate
	 * way to express if/else logic.
	 * Here we also show the 'case" expressions can
	 * be non-contante
	 * ==================================================
	 */
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	fmt.Println("===================================")

	/**
	 * ==================================================
	 * A type 'switch' compare types instead of values.
	 * You can use this to discover the type of an
	 * interface value. In this example, the variable
	 * t will have type corresponding to its clause.
	 * ==================================================
	 */
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
