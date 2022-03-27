package main

import "fmt"

func main() {
	fmt.Println(soma(10, 20))
	fmt.Println("test")

	bi := [][]int{
		{0, 1, 2, 3},
		{4, 5, 6, 7},
		{8, 9, 10, 11},
	}

	fmt.Println(bi[1][1])
}

func soma(a int, b int) int {
	return a + b
}
