package main

import "fmt"

// 定義 Line 和 Point 結構，讓主程式能印出預期的結果
func main() {
	line := Line{Point{0, 0}, Point{3, 4}}
	fmt.Println(line.P1.X) // 印出 0
	fmt.Println(line.P2.Y) // 印出 4
}

// My Code
type Line struct {
	P1 Point
	P2 Point
}

type Point struct {
	X int
	Y int
}
