package main

import "fmt"

// 承上，設計 Animal 介面，滿足主程式需求
func main() {
	m := Man{"John", 30}
	// 修正以下參數傳遞的方式，需印出：I’m John, 31 years old.
	birthdayTalk(&m)
	b := Bird{"KiKi", 2, 5}
	// 修正以下參數傳遞的方式，需印出：I’m bird KiKi, 3 years old. Fly 5 miles a day.
	birthdayTalk(&b)
}

func birthdayTalk(a Animal) {
	a.Grow(1)
	a.Talk()
}

// My Code
type Animal interface {
	Grow(age int)
	Talk()
}

type Man struct {
	name string
	age  int
}

func (m *Man) Grow(age int) {
	m.age += age
}

func (m Man) Talk() {
	fmt.Printf("I’m %s, %d years old\n", m.name, m.age)
}

type Bird struct {
	name        string
	age         int
	flyDistance int
}

func (m *Bird) Grow(age int) {
	m.age += age
}

func (b Bird) Talk() {
	fmt.Printf("I’m bird %s, %d years old. Fly %d miles a day.\n", b.name, b.age, b.flyDistance)
}
