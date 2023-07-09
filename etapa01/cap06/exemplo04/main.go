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
	add := func(a, b int) int {
		return a + b
	}
	fmt.Println(add(1, 100))
}
