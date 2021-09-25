package main

import "fmt"

func main() {
	printToMax(5)  // 印出 0, 1, 2, 3, 4, 5
	printToMax(15) // 印出 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10
}
func printToMax(max int) {
	// 使用 for 迴圈從 0 印到 max，但若 max 超過 10，只印到 10 為止
	// 可以假設 max 一定會大於 0

	// MyCode
	for i := 0; i <= 10; i++ {
		if i == max || i == 10 {
			fmt.Println(i)
			break
		} else {
			fmt.Printf("%d, ", i)
			continue
		}
	}
}
