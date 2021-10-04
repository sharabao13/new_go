package main

import "fmt"

func main() {
	// 字符串
	// ""  可解释  \转义
	// `` 原生字符串

	var s1 string = "hcx"
	var s2 string = `hcx`
	var s3 string = `This is a "book",This is a pen`

	fmt.Println(s1, s2, s3)

	//字符串操作
	// 算术运算  +
	fmt.Println("a" + "b")

	//比较运算
	fmt.Println("ab" == "a")
	fmt.Println("ab" != "a")
	fmt.Println("ab" > "a")
	fmt.Println("ab" >= "a")
	fmt.Println("ab" < "a")
	fmt.Println("ab" <= "a")

	s4 := "i am"
	s4 += " bao"
	fmt.Println(s4)

	//字符串索引 [0-n-1]
	s5 := "Test"
	fmt.Println(s5[0])

	//切片[start:end-1]
	s6 := "abcdefg"
	fmt.Println(s6[1:3])

	// 长度 len
	fmt.Println(len(s6))

}
