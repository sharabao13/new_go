package main

import "fmt"

func main() {
	// 指针 *int *string *..
	// 指针赋值 &
	//取指针 *

	i1 := 3
	i2 := i1
	i2 = 4
	fmt.Println(i1, i2)

	//指针
	var p1 *int
	p1 = &i1
	fmt.Println(*p1)

	s1 := "abc"
	var p2 *string = &s1
	fmt.Println(*p2)
	fmt.Printf("p2=%T,p2=%p\n", p2, p2)
}
