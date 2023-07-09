package main

import "fmt"

func main() {

	defer func() {
		fmt.Println("Starting")
		str := recover()
		fmt.Println(str)
	}()

	panic("Panic!")
}
