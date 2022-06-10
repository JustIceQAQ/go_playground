package main

import (
	"fmt"
	"math/rand"
	"time"
)

func timeCost() func() {
	start := time.Now()
	return func() {
		tc := time.Since(start)
		fmt.Printf("本次遊玩耗時: %v\n", tc)
	}

}

func main() {

	defer timeCost()()
	rand.Seed(time.Now().UnixNano())
	ansNumber := rand.Intn(100)
	//fmt.Print("目前正確答為: ", ansNumber, "\n")
	var userInputNumber int
	userAnswerCount := 1

	var minRangeNumber, maxRangeNumber = 0, 100

	fmt.Printf("請輸入你的答案(%d ~ %d): ", minRangeNumber, maxRangeNumber)
	for ansNumber != userInputNumber {
		fmt.Scanln(&userInputNumber)
		if ansNumber != userInputNumber {
			// 正確答案 < 使用者輸入答案 < 目前最大答案
			if ansNumber < userInputNumber && userInputNumber <= maxRangeNumber {
				maxRangeNumber = userInputNumber
			}
			// 目前最小答案<  使用者輸入答案 < 正確答案
			if minRangeNumber < userInputNumber && userInputNumber < ansNumber {
				minRangeNumber = userInputNumber
			}
			fmt.Printf("猜錯了! 再猜一次 (%d ~ %d):", minRangeNumber, maxRangeNumber)
			userAnswerCount += 1
		}
	}
	fmt.Println("恭喜你答對了，答案就是: ", userInputNumber)
	fmt.Println("你總共猜了 ", userAnswerCount, " 次")
}
