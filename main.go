package main

import (
	"fmt"
)

func Calculate(x int) (result int) {
	result = x + 2
	return result
}

func divide(x float64, y float64) (result float64) {
	result = x / y
	return result
}

func main() {
	fmt.Println("Adding 2 to an integer")
	result := divide(10.0, 0.0)
	fmt.Println(result)
}
