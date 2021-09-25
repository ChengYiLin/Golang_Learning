package main

import "fmt"

func main() {
	n := 3
	add(&n, 2)
	fmt.Println(n) // 印出 5
	add(&n, 5)
	fmt.Println(n) // 印出 10
}

// 完成一個 add 函式，讓主程式能印出預期的結果

// MyCode
func add(address *int, val int) {
	*address = *address + val
}
