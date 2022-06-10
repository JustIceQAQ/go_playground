package main

import (
	"fmt"
)
import UserInputStruct "go_playground/exp_calculator_utils"

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

func doAllCalculator(userInput UserInputStruct.UserInput) {
	fmt.Println(formatPrint("+", addition(userInput.A, userInput.B)))
	fmt.Println(formatPrint("-", subtraction(userInput.A, userInput.B)))
	fmt.Println(formatPrint("*", multiplication(userInput.A, userInput.B)))
	fmt.Println(formatPrint("/", division(userInput.A, userInput.B)))
}

func main() {
	userInput := UserInputStruct.UserInput{}
	fmt.Println("Calculator")
	fmt.Println("Pls input a and b value: ")
	fmt.Scanln(&userInput.A, &userInput.B) // 50 5
	doAllCalculator(userInput)
}
