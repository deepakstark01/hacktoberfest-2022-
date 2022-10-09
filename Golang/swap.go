package main

import "fmt"

func main() {
	var num1 int = 11
	var num2 int = 22

	fmt.Println("Numbers before swapping:")
	fmt.Println("Num1: ", num1)
	fmt.Println("Num2: ", num2)

	swap(&num1, &num2)

	fmt.Println("Numbers after swapping:")
	fmt.Println("Num1: ", num1)
	fmt.Println("Num2: ", num2)
}

// swap takes two pointers of numbers and swap those by arithmetic operations
func swap(num1, num2 *int) {
	*num1 += *num2
	*num2 = *num1 - *num2
	*num1 -= *num2
}
