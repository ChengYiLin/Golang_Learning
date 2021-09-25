// Question : 請建立 divide 函式，回傳 n1 除以 n2 的結果，結果必須是小數 ( 浮點數 ) 的資料型態
package main

import "fmt"

func main() {
	var n1 int = 2
	var n2 int = 3
	result := divide(n1, n2)
	fmt.Println(result) // 此行程式必須輸出 n1 除以 n2 的結果
}

// My Code
func divide(n1 int, n2 int) float32 {
	var Fn1 float32 = float32(n1)
	var Fn2 float32 = float32(n2)

	return (Fn1 / Fn2)
}
