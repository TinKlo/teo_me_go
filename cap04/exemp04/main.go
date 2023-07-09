package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		var check string
		if i%2 == 0 {
			check = "par"
		} else {
			check = "Ímpar"
		}
		fmt.Println(i, check)

	}
}
