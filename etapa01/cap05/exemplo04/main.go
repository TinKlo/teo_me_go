package main

import "fmt"

func main() {
	var x [5]int
	x[0] = 98
	x[1] = 10
	x[2] = 40
	x[3] = 12
	x[4] = 38
	fmt.Println(x)
	total := 0
	for _, value := range x {
		total += value

	}
	fmt.Println(total / len(x))
}
