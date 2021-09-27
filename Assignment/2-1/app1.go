package main

import "fmt"

func main() {
	data1 := [3]int{3, 1, 2}
	calculate(data1[:]) // 請修正參數的處理方式，讓函式能印出 3+1+2 的結果 6
	data2 := [4]int{10, 3, -5, 1}
	calculate(data2[:]) // 請修正參數的處理方式，讓函式能印出 10+3+(-5)+1 的結果 9
}

// 請建立一個 calculate 函式，接受 int 型態切片，計算並印出原始陣列中所有整數的總和

// My Code
func calculate(vals []int) {
	var result int = 0
	for i := 0; i < len(vals); i++ {
		result += vals[i]
	}

	fmt.Println(result)
}
