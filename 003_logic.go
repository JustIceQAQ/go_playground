package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	fmt.Printf("003_logic.go")
	// random seed set
	rand.Seed(time.Now().UnixNano())
	var a = rand.Intn(40)

	// number if else
	var q = a%5 == 0
	if q {
		strFinal := fmt.Sprintln("Yes", a, "%%", 5, "is 0")
		fmt.Printf(strFinal)
	} else {
		strFinal := fmt.Sprintln("No", a, "%%", 5, "is not 0")
		fmt.Printf(strFinal)
	}

	// string if else
	var s1 = "QAQ"
	if strings.Contains(s1, "Q") {
		fmt.Printf("OK")
	}

	// switch

	var aa = 10

	switch aa {
	case 1:
		fmt.Println("aa is 1")
	case 10:
		fmt.Println("aa is 10")
	default:
		fmt.Println("aa is 10")
	}

}
