package main

import "fmt"

func main() {
	calculate(1, 3) // 函式要能夠計算 1+2+3，最後印出 6
	calculate(4, 8) // 函式要能夠計算 4+5+6+7+8，最後印出 30
}
func calculate(min int, max int) {
	// 請運用 for 迴圈，計算並印出 min+(min+1)+(min+2)+...+max 的總和
	// 可以假設 max 一定會大於 min

	// My Code
	var result int = 0

	for i := min; i <= max; i++ {
		result += i
	}

	fmt.Println(result)
}
