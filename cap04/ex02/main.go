//write a every program between 1 and 100 that are divisible by 3 and integers

package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		// fmt.Println(i)
		if i%3 == 0 {
			fmt.Println(i)
		}
	}

}
