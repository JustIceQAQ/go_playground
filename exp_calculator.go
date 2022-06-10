package main

import (
	"fmt"
)
import UserInputStruct "go_playground/exp_calculator_utils"

func main() {
	userInput := UserInputStruct.UserInput{}
	fmt.Println("Calculator")
	fmt.Println("Pls input a and b value: ")
	_, err := fmt.Scanln(&userInput.A, &userInput.B)
	if err != nil {
		return
	} // 50 5
	userInput.ShowValue()
	userInput.AllCalculator()
}
