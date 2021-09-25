package main

import "fmt"

// 定義一個 Man 結構，讓主程式能印出預期的結果
func main() {
	m := Man{"小明", 30}
	fmt.Println(m.Name) // 印出 小明
	fmt.Println(m.Age)  // 印出 30
}

// My Code
type Man struct {
	Name string
	Age  int
}
