package main

import "fmt"

func main() {
	fmt.Println(Soma(11, 10))
}

func Soma(a int, b int) int {
	return a + b
}

func sub(a int, b int) int {
	return a - b
}

func times(a int, b int) int {
	return a * b
}

func div(a int, b int) int {
	return a / b
}

func pow(a int, b int) int {
	result := a

	for i := 1; i < b; i++ {
		result *= a
	}

	return result
}
