// Question : 請輸出 n1 除以 n2 的結果，結果必須是小數 ( 浮點數 ) 的資料型態
package main

import "fmt"

func main() {
	n1 := 2
	n2 := 3

	// My Code
	var Fn1 float32 = float32(n1)
	var Fn2 float32 = float32(n2)

	fmt.Println(Fn1 / Fn2)
}
