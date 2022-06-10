package exp_calculator_utils

import "fmt"

type UserInput struct {
	A int // first value
	B int // second value
}

func (u *UserInput) ShowValue() {
	// format print info
	fmt.Println(fmt.Sprint("A = ", u.A, ", B = ", u.B))
}

func (u *UserInput) Addition() int {
	return u.A + u.B
}
func (u *UserInput) Subtraction() int {
	return u.A - u.B
}
func (u *UserInput) Multiplication() int {
	return u.A * u.B
}
func (u *UserInput) Division() int {
	return u.A / u.B
}

func (u *UserInput) FormatPrint(icon string, r int) string {
	return fmt.Sprint("a ", icon, " b = ", r)
}

func (u *UserInput) AllCalculator() {
	fmt.Println(u.FormatPrint("+", u.Addition()))
	fmt.Println(u.FormatPrint("-", u.Subtraction()))
	fmt.Println(u.FormatPrint("*", u.Multiplication()))
	fmt.Println(u.FormatPrint("/", u.Division()))
}
