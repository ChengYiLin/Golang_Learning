package main

import "fmt"

// 承上題，定義一個 Man 結構，讓主程式能印出預期的結果
func main() {
	m := Man{"小明", 20}
	addAge(&m, 1)
	fmt.Println(m.Name) // 印出 小明
	fmt.Println(m.Age)  // 印出 21
}

// 完成一個 addAge 函式，讓主程式能印出預期的結果

// My Code
type Man struct {
	Name string
	Age  int
}

func addAge(man *Man, age int) {
	man.Age = man.Age + age
}
