package main

import (
	"fmt"

	"./libs/tree"
)

func walk(t *tree.Tree, ch chan int) {
	recWalk(t, ch)
	close(ch)
}

func recWalk(t *tree.Tree, ch chan int) {
	if t != nil {
		recWalk(t.Left, ch)
		ch <- t.Value
		recWalk(t.Right, ch)
	}
}

func same(t1 *tree.Tree, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go walk(t1, ch1)
	go walk(t2, ch2)

	for {
		x1, ok1 := <-ch1
		x2, ok2 := <-ch2
		switch {
		case ok1 != ok2:
			return false
		case !ok1 && !ok2:
			return true
		case x1 != x2:
			return false
		default:
		}
	}
}

func main() {
	ch := make(chan int)
	go walk(tree.New(1), ch)
	for v := range ch {
		fmt.Println(v)
	}
	fmt.Println(same(tree.New(1), tree.New(1)))
	fmt.Println(same(tree.New(1), tree.New(2)))
}
