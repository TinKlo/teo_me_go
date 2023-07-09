package main

import "fmt"

func constructorPairGen() func() int {
	i := 0
	return func() (ret int) {
		ret = i
		i += 2
		return
	}
}

func main() {
	fmt.Println("Begin")
	proximoPar := constructorPairGen()
	fmt.Println(proximoPar())
	fmt.Println(proximoPar())
	fmt.Println(proximoPar())
}
