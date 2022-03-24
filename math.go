package main

import "fmt"

func main() {
	fmt.Println(Soma(10, 20))
	fmt.Println("test")
}

func Soma(a int, b int) int {
	return a + b
}
