package main

import "fmt"

func main() {
	slice := make([]int, 3, 9)
	fmt.Println(slice)
	fmt.Println("O tamanho da fatia e: ", len(slice))
}
