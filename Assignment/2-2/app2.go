package main

import (
	"fmt"
)

func main() {
	// My Code
	var data = make(map[string]string, 3)

	// 在此，請適當的定義 data 變數，存放地圖資料，讓以下程式能正常運作
	data["apple"] = "蘋果"
	data["dog"] = "狗狗"
	data["cat"] = "貓貓"
	// 在此，請要求使用者透過終端機輸入英文字，利用 data 進行對應中文字的查詢
	// 若輸入的英文字包含在 data 的 key 中，則印出對應的中文；若無，印出沒有查到

	// My Code
	var inputTxt string
	fmt.Printf("Enter Text : ")
	fmt.Scanf("%s", &inputTxt)

	outputTxt, ok := data[string(inputTxt)]
	if !ok {
		fmt.Println("沒有查到")
		return
	}
	fmt.Println(outputTxt)
}
