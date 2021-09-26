package main

import "fmt"

func main() {
	// 在此，請適當的定義 data 變數，存放雙層切片的資料，讓以下程式能正常運作
	// My Code
	var data = make([][]int, 2)
	data[0] = make([]int, 3)
	data[1] = make([]int, 3)

	data[0][0] = 10
	data[0][1] = 5
	data[0][2] = 1
	data[1][0] = 3
	data[1][1] = 8
	data[1][2] = 4

	// 在此，請運用雙層 for 迴圈，計算以上六個數字的總和，並印出結果
	// My Code
	var result int = 0
	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data[i]); j++ {
			result += data[i][j]
		}
	}

	fmt.Println(result)
}
