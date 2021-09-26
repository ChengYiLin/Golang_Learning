package main

import "fmt"

// 承上，另外替 Man 結構加入 Grow 方法用來表達年齡增長，並滿足主程式的需求
func main() {
	m := Man{"John", 30}
	m.Grow(2)
	m.Talk() // 印出：I’m John, 32 years old
	m = Man{"Bob", 18}
	m.Grow(1)
	m.Talk() // 印出：I’m Bob, 19 years old
}

// My Code
// Ref : https://golang.org/doc/faq#methods_on_values_or_pointers
type Man struct {
	Name string
	Age  int
}

func (m Man) Talk() {
	fmt.Printf("I’m %s, %d years old\n", m.Name, m.Age)
}

func (m *Man) Grow(age int) {
	m.Age += age
}
