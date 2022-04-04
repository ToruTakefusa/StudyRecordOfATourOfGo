package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println("first")
	fmt.Println(<-ch)
	fmt.Println("second")
	fmt.Println(<-ch)
}
