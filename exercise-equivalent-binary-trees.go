package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	var walk func(t *tree.Tree, ch chan int)
	walk = func(t *tree.Tree, ch chan int) {
		if t == nil {
			return
		}
		if t.Left != nil {
			walk(t.Left, ch)
		}
		ch <- t.Value
		if t.Right != nil {
			walk(t.Right, ch)
		}
	}
	walk(t, ch)
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for i := 0; i < 10; i++ {
		if <-ch1 == <-ch2 {
			continue
		} else {
			return false
		}
	}
	return true
}

func main() {
	ch := make(chan int, 10)
	go Walk(tree.New(1), ch)
	for val := range ch {
		fmt.Print(val, ",")
	}
	fmt.Println()

	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
