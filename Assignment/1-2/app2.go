package main

import "fmt"

func main() {
	n := 3
	n = add(n, 2)
	fmt.Println(n) // 印出 5
	n = add(n, 5)
	fmt.Println(n) // 印出 10
}

// 完成一個 add 函式，讓主程式能印出預期的結果

// My Code
func add(val1 int, val2 int) int {
	return val1 + val2
}
