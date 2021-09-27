package main

import (
	"fmt"
)

// 請設計 Man 結構，定義姓名和年齡，並滿足主程式的需求
func main() {
	m := Man{"John", 30}
	print(m) // 印出：I’m John, 30 years old.
	s := "Hello"
	print(s) // 印出：String is Hello
	f := 3.5
	print(f) // 印出：Float is 3.5
}

// 請運用空白介面當參數，建立 print 方法，滿足主程式需求

// My Code
type Man struct {
	name string
	age  int
}

type Printer interface{}

func print(param interface{}) {
	switch param.(type) {
	default:
		fmt.Printf("unexpected type %T\n", param)

	case Man:
		m := param.(Man)
		fmt.Printf("I’m %s, %d years old.\n", m.name, m.age)

	case string:
		fmt.Printf("String is  %T\n", param)

	case float64:
		fmt.Printf("Float is %T\n", param)
	}
}
