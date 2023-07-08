package main

import "fmt"

func main() {
	fmt.Println("Entre com o valor a ser convertido de celsius para Farenheit: ")
	var f float64
	fmt.Scanf("%f", &f)

	c := (f - 32) * (5. / 9.)
	fmt.Println(c)
}
