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
	for i := 0; i < 5; i++ {
		total += x[i]

	}
	fmt.Println(total / 5)
}
