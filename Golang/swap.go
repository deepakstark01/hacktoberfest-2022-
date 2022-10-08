package main

import "fmt"

func main() {
	var num1 int = 11
	var num2 int = 22
	var num3 int = 0

	fmt.Println("Numbers before swapping:")
	fmt.Println("Num1: ", num1)
	fmt.Println("Num2: ", num2)

	num3 = num1
	num1 = num2
	num2 = num3

	fmt.Println("Numbers after swapping:")
	fmt.Println("Num1: ", num1)
	fmt.Println("Num2: ", num2)
}
