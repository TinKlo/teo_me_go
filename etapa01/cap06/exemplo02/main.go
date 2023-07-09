package main

import "fmt"

func aveg(x []float64) float64 {
	fmt.Println("Begin")
	total := 0.0
	for _, v := range x {
		total += v
	}
	return total / float64(len(x))
}

func main() {
	x := []float64{1, 2, 3, 4, 5}

	fmt.Println("Avg of x: ", aveg(x))
}
