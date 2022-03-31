package main

import "fmt"

func fibonacci() func() int {
	now := 0
	before := 0
	count := 0

	return func() int {
		if count == 0 {
			count++
			return now
		} else if count == 1 {
			count++
			now = 1
			return now
		} else {
			tmp := now
			now += before
			before = tmp
			count++
			return now
		}
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
