package main

import "fmt"

func main() {
	/*
			数值类型
		    整形 int
		    无符号整形 uint
		    占位符 %d
		    字面量 十进制 八进制 十六进制
	*/
	i1 := 18
	fmt.Printf("i1=%T,i1=%d\n", i1, i1)

	fmt.Println(0777)
	fmt.Println(10)

	//操作
	// 算术运算  + - * / %

	fmt.Println("算术运算")
	fmt.Println(1 + 2)
	fmt.Println(1 - 2)
	fmt.Println(1 * 2)
	fmt.Println(9 / 2)
	fmt.Println(9 % 2)
	fmt.Println("-----------------------")

	//自增 ++  自减 --
	age := 10
	age++
	fmt.Println(age)
	age--
	fmt.Println(age)

	var a1 byte = 'A'
	var a2 rune = '中'
	fmt.Println(a1, a2)

}
