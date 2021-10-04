package main

import (
	"fmt"
)

func main() {
	/*
				一维数组: 存储多个数据数值的本身
			 		a1 := [3]int{1,2,3}
				二维数组： 存储的是一维的一维
					a2 := [3][4]{{},{},{}}
					该二维的数组的长度，就是3
					存储的元素是一维数组的，一维数组的元素是数值，每个一维数组的长度为4

		        多维数组
	*/
	a2 := [3][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 13}}
	fmt.Println(a2)
	fmt.Println(a2[0][1])
	fmt.Println("-------------------")

	//二维数组的遍历 for range
	for _, v1 := range a2 {
		for _, v2 := range v1 {
			fmt.Println(v2)
		}
	}

	//多维数组
	a4 := [1][2][2][1]int{{{{1}, {2}}, {{3}, {4}}}}
	fmt.Printf("%T\n", a4)
	fmt.Println(a4)
	fmt.Println("-------------")
	for _, v1 := range a4 {
		for _, v2 := range v1 {
			for _, v3 := range v2 {
				for _, v4 := range v3 {
					fmt.Println(v4)
				}
			}
		}
	}

	//五维数组
	a5 := [2][3][1][4][2]int{{{{{1, 2}, {3, 4}, {5, 6}, {7, 8}}}, {{{9, 10}, {11, 12}, {13, 14}, {15, 16}}}, {{{17, 18}, {19, 20}, {21, 22}, {23, 24}}}}, {{{{25, 26}, {27, 28}, {29, 30}, {31, 32}}}, {{{33, 34}, {35, 36}, {37, 38}, {39, 40}}}, {{{41, 42}, {43, 44}, {45, 46}, {47, 48}}}}}
	fmt.Println(a5)
	fmt.Println("-------------")
	for _, v1 := range a5 {
		for _, v2 := range v1 {
			for _, v3 := range v2 {
				for _, v4 := range v3 {
					for _, v5 := range v4 {
						fmt.Println(v5)
					}
				}
			}
		}
	}

}
