package main

import "fmt"

func main() {
	/*
			可变参数
				go函数支持变参，接收变参的函数是有着不定量的参数，为了做到这点，首先需要定义函数使其接收可变参
			概念: 一个函数的参数类型确定，个数不确定，就可以使用可变参
			语法:
				func funName(可变参数名 ... type){}
				可变参数返回的是slice
		    注意：
				1. 如果一个函数是可变参数，同时还有其他参数，可以变参数要放在末尾
				2. 一个函数的可变参数最多只能有一个
	*/
	//求和
	getSums(1, 2)

	//可变参数传入slice 相当于append （slice...）
	s1 := []int{1, 2, 3}
	getSums(s1...)
}

func getSums(nums ...int) {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	fmt.Println("总和是: ", sum)
}
