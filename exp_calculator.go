package main

import "fmt"

var userInputA, userInputB int

func addition(a int, b int) int {
	return a + b
}
func subtraction(a int, b int) int {
	return a - b
}
func multiplication(a int, b int) int {
	return a * b
}

func division(a int, b int) int {
	return a / b
}

func formatPrint(icon string, r int) string {
	return fmt.Sprint("a ", icon, " b = ", r)
}

func doAllCalculator(a int, b int) {
	fmt.Println(formatPrint("+", addition(a, b)))
	fmt.Println(formatPrint("-", subtraction(a, b)))
	fmt.Println(formatPrint("*", multiplication(a, b)))
	fmt.Println(formatPrint("/", division(a, b)))
}

func main() {
	fmt.Println("Calculator")
	fmt.Println("Pls input a and b value: ")
	fmt.Scanln(&userInputA, &userInputB) // 50 5
	doAllCalculator(userInputA, userInputB)
}
