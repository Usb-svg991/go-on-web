package main

import "fmt"

func square(x int) int {
	return x * x
}

func main() {
	nums := []int{2, 4, 6, 8, 10}
	ch := make(chan int)

	for _, n := range nums {
		go func(x int) {
			ch <- square(x)
		}(n)
	}

	for range nums {
		fmt.Println(<-ch)
	}
}
