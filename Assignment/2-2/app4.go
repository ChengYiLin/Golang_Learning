package main

import (
	"errors"
	"fmt"
)

func main() {
	n1 := 3
	var n2 int
	fmt.Printf("Enter a Number: ")
	fmt.Scanf("%d", &n2)
	result, err := divide(n1, n2)
	// 在此加入判斷式
	// 如果 err 是 nil，印出錯誤訊息 Cannot Divide by Zero，否則印出除法的結果

	// My Code
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}

// 請建立 divide 方法
// 如果第二個參數 ( 除數 ) 是 0：傳回結果 0，以及錯誤訊息為 Cannot Divide by Zero
// 否則：傳回除法的結果，以及錯誤訊息為 nil

// My Code
func divide(n1 int, n2 int) (float64, error) {
	if n2 == 0 {
		return 0, errors.New("Cannot Divide by Zero")
	}

	return float64(n1) / float64(n2), nil
}
