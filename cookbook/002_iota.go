package main

import "fmt"

const (
	A = iota
	B
	C
	D
	E = iota * 0.1
	F
	G
	H
	I = iota + 100
)

func main() {
	a := ""
	fmt.Scanf("%s", &a)
	fmt.Println(A, B, C, D, E, F, G, H, I, a)
}
