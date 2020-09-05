package main

import "fmt"

/**
 * ============================================================
 * This person struct type has name and age fields.
 * ============================================================
 */
type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	/**
	 * ============================================================
	 * You can safely return a pointer to local variable as a local
	 * variable will survive the scope of the funtion.
	 * ============================================================
	 */
	p := person{name: name}
	p.age = 42
	return &p
}

func main() {
	/**
	 * ============================================================
	 * This syntax creates a new struct.
	 * ============================================================
	 */
	fmt.Println(person{"Bob", 42})

	fmt.Println("=================================================")

	/**
	 * ============================================================
	 * You can name the field when initializing a struct.
	 * ============================================================
	 */
	fmt.Println(person{name: "Alice", age: 30})

	fmt.Println("=================================================")

	/**
	 * ============================================================
	 * Omitted field will be zero-value.
	 * ============================================================
	 */
	fmt.Println(person{name: "Fred"})

	fmt.Println("=================================================")

	/**
	 * ============================================================
	 * An & prefix yields a pointer to the struct.
	 * ============================================================
	 */
	fmt.Println(&person{name: "Ann", age: 40})

	fmt.Println("=================================================")

	/**
	 * ============================================================
	 * It's idiomatic to encapsulate a new struct creation in
	 * constructor functions.
	 * ============================================================
	 */
	fmt.Println(newPerson("John"))

	fmt.Println("=================================================")

	/**
	 * ============================================================
	 * Access struct fields with a dot.
	 * ============================================================
	 */
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	fmt.Println("=================================================")

	/**
	 * ============================================================
	 * You can also use dots with struct pointers - The pointers
	 * are automatically derefenced.
	 *
	 * Struct are mutable.
	 * ============================================================
	 */
	sp := &s
	fmt.Println(sp.age)
	sp.age = 51
	fmt.Println(sp.age)
}
