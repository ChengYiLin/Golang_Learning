package main

import "fmt"

// 請設計 Man 結構，定義姓名和年齡，並且替 Man 結構加入 Talk 方法，滿足主程式的需求
func main() {
	m := Man{"John", 30}
	m.Talk() // 印出：I’m John, 30 years old
	m = Man{"Bob", 18}
	m.Talk() // 印出：I’m Bob, 18 years old
}

// My Code
type Man struct {
	Name string
	Age  int
}

func (m Man) Talk() {
	fmt.Printf("I’m %s, %d years old\n", m.Name, m.Age)
}
