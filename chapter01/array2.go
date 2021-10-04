package main

import "fmt"

func main() {
	/*
				数据类型
					基本类型: int,float,string bool
			        复合类型: array slice map function pointer channel

				值类型: 理解为存储数组本身
						int float string bool array
		    	引用类型: 理解为存储的数据的内存地址
	*/

	//数组排序 冒泡排序
	arr := [5]int{15, 23, 8, 10, 7}
	for i := 1; i < len(arr); i++ {
		for j := 0; j < len(arr)-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
		fmt.Println(arr)
	}
}
