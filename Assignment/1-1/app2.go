// Question : 請建立 divide 函式，輸出 n1 除以 n2 的結果，結果必須是小數 ( 浮點數 ) 的資料型態
package main

import "fmt"

func main() {
	var n1 int = 2
	var n2 int = 3
	divide(n1, n2)
}

// My Code
func divide(n1 int, n2 int) {
	var Fn1 float32 = float32(n1)
	var Fn2 float32 = float32(n2)

	fmt.Println(Fn1 / Fn2)
}
