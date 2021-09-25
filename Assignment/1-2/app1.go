package main

import "fmt"

func main() {
	n := 3
	nPtr := &n
	// 修正以下程式，讓程式能印出原來的數字 3
	fmt.Println("Address : ", nPtr)

	// MyCode
	fmt.Println("Value : ", *nPtr)
}
