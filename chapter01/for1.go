package main

import "fmt"

func main() {
	/*
				for 循环嵌套
			    题目1
			    打印如下图形
			     *****
				 *****
				 *****
				 *****
				 *****

		        题目2
		        打印99乘法表
	*/
	//题目1
	for i := 1; i <= 5; i++ {
		for j := 1; j <= 5; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	fmt.Println("---------------------------")

	//题目2
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d x %d = %d\t", j, i, j*i)
		}
		fmt.Println()
		//fmt.Printf("\t")
	}
	//fmt.Println()

}
