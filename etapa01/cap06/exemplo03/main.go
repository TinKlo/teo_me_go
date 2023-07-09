package main

import "fmt"

func summ(args ...int) int {
	fmt.Println("Begin")
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}

func main() {
	x := []int{1, 2, 3, 4, 5}

	fmt.Println("Sum of x: ", summ(x...))
}
