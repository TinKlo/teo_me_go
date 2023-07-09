package main

import "fmt"

func main() {
	fmt.Println("Como podemos acessar o quarto elemento de um array de uma fatia")
	slice1 := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(slice1)
	i := slice1[3]
	fmt.Println(i)
}
