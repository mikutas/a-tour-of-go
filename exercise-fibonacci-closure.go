package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	index := 0
	a := 1
	b := 0
	return func() int {
		if index == 0 {
			index++
			return 0
		} else if index == 1 {
			index++
			return 1
		} else {
			ret := a + b
			b = a
			a = ret
			index++
			return ret
		}
		return 0
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
