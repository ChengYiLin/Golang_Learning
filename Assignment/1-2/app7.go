package main

import "fmt"

func main() {
	data := [10]int{3, 10, 2, -5, 2, 10, 1, -2, 8, 15}
	// 請在此撰寫一段程式，使用 for 迴圈與 len() 內建函式搭配
	// 計算並印出陣列中數字的總和 44

	//  My Code
	var total int = 0
	for i := 0; i < len(data); i++ {
		total += data[i]
	}

	fmt.Println(total)
}
