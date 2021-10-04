package main

import "fmt"

func main() {
	/*
	 	常量定义
	     特性 不可更改
	*/
	const s1 string = "hcx"
	//s1 = "hfb" 不可修改
	fmt.Println(s1)

	const (
		s2 string = "hcx"
		i2 int    = 18
	)
	fmt.Println(s2, i2)

	const s3, s4 string = "hcx", "gfb"
	fmt.Println(s3, s4)

	const s5, i5 = "hxc", 18
	fmt.Println(s5, i5)

	const (
		i6 = 7
		i7
		i8
		s9 = "hcx"
		s10
	)
	fmt.Println(i6, i7, i8, s9, s10)
	//常量枚举类型 iota
	const (
		a1 int = iota
		a2
		a3
		a4
	)
	fmt.Println(a1, a2, a3, a4)

	const (
		b1 = iota
		b2
		b3
		b4 = iota + 1
		b5
	)
	fmt.Println(b1, b2, b3, b4, b5)

	const (
		c1 = (iota + 1) * 100
		c2
		c3
	)
	fmt.Println(c1, c2, c3)
}
