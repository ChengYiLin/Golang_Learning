package main

import "fmt"

func main() {
	n1 := 3
	n2 := 4
	r1, r2 := calculate(n1, n2)
	fmt.Println(r1, r2) // 這裡要印出乘法的結果 13，除法的結果 0.75
}

// 請建立 calculate 方法，同時回傳乘法和除法的結果，滿足主程式需求

// My Code
func calculate(n1 int, n2 int) (int, float64) {
	multiple := n1 * n2
	division := float64(n1) / float64(n2)

	return multiple, division
}
