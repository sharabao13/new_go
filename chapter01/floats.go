package main

import "fmt"

func main() {
	var height float64
	height = 180.67
	fmt.Println(height)

	var t1 float64 = 67.88
	fmt.Printf("t1=%T,t1=%.2f\n", t1, t1)

	// 算术运算
	fmt.Println(1.22 + 2.09)
	fmt.Println(1.22 - 2.09)
	fmt.Println(1.22 * 2.09)
	fmt.Println(1.22 / 2.09)
}
