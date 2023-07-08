package main

import "fmt"

func main() {
	const x string = "Olá Mundo!"
	fmt.Println(x)
	fmt.Println("Entre com um valor: ")
	var a int
	fmt.Scanf("%d", &a) //should be explored further in next chapters

	res := a * 2
	fmt.Println(res)
}
