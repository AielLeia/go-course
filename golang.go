package main

import "fmt"

func main() {
	/**
	 * ======================================================
	 * Maps are Go's built-in associative data type
	 * ======================================================
	 */

	/**
	 * ======================================================
	 * To create an empty mapn use the builtin 'make':
	 * 'make(map[key-type]val-type)'.
	 * ======================================================
	 */
	m := make(map[string]int)

	fmt.Println("=======================================================")

	/**
	 * ======================================================
	 * Set key/value pairs using typical 'name[key] = val'
	 * syntax.
	 * ======================================================
	 */
	m["k1"] = 7
	m["k2"] = 13
	fmt.Println("map:", m)

	fmt.Println("=======================================================")

	/**
	 * ======================================================
	 * get a value for a key with 'name[key]'
	 * ======================================================
	 */
	v1 := m["k1"]
	fmt.Println("v1:", v1)

	fmt.Println("=======================================================")

	/**
	 * ======================================================
	 * The builtin 'len' returns the number of key/value
	 * pairs when called on a map.
	 * ======================================================
	 */
	fmt.Println("len:", len(m))

	fmt.Println("=======================================================")

	/**
	 * ======================================================
	 * The builtin 'delete' removes key/value pairs from a
	 * map.
	 * ======================================================
	 */
	delete(m, "k2")
	fmt.Println("map:", m)

	fmt.Println("=======================================================")

	/**
	 * ======================================================
	 * The optinal second return value when getting a value
	 * from a map indicate if the key was present in map.
	 * This can be used to disambiguate between missin keys
	 * and keys with zero values like '0' from '""'. Here
	 * we didn't need the value itself, so we ignored it with
	 * the blank identifier _.
	 * ======================================================
	 */
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	fmt.Println("=======================================================")

	/**
	 * ======================================================
	 * The optinal second return value when getting a value
	 * from a map indicate if the key was present in map.
	 * This can be used to disambiguate between missin keys
	 * and keys with zero values like '0' from '""'. Here
	 * we didn't need the value itself, so we ignored it with
	 * the blank identifier _.
	 * ======================================================
	 */
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}
