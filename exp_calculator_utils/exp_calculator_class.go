package exp_calculator_utils

import "fmt"

type UserInput struct {
	A int
	B int
}

func (u *UserInput) ShowValue() {
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
