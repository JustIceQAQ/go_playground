package main

import (
	"fmt"
)
import UserInputStruct "go_playground/exp_calculator_utils"

func formatPrint(icon string, r int) string {
	return fmt.Sprint("a ", icon, " b = ", r)
}

func doAllCalculator(userInput UserInputStruct.UserInput) {
	fmt.Println(formatPrint("+", userInput.Addition()))
	fmt.Println(formatPrint("-", userInput.Subtraction()))
	fmt.Println(formatPrint("*", userInput.Multiplication()))
	fmt.Println(formatPrint("/", userInput.Division()))
}

func main() {
	userInput := UserInputStruct.UserInput{}
	fmt.Println("Calculator")
	fmt.Println("Pls input a and b value: ")
	_, err := fmt.Scanln(&userInput.A, &userInput.B)
	if err == nil {

		return
	} // 50 5
	userInput.ShowValue()
	doAllCalculator(userInput)
}
