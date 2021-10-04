package main

import "fmt"

func main() {
	/*
		参数传递
			1. 值传递，传递的是参数的副本。修改数据对于原始数据没有影响
				值类型的数据都是值传递
			2. 引用传递
	*/
	//数组值类型的数据
	arr1 := [4]int{1, 2, 3, 4}
	fmt.Println("函数调用前数组的数据", arr1)
	fun2(arr1)
	fmt.Println("函数调用后数组的数据", arr1)
	fmt.Println("-------------------------")

	//切片引用类型的数据
	s1 := []int{1, 2, 3, 4}
	fmt.Println("函数调用前切片的数据", s1)
	fun3(s1)
	fmt.Println("函数调用后切片的数据", s1)

}

func fun2(arr2 [4]int) {
	fmt.Println("函数中数组的数据", arr2)
	arr2[0] = 100
	fmt.Println("函数中修改后数组的数据", arr2)
}

func fun3(s2 []int) {
	fmt.Println("函数中切片的数据", s2)
	s2[0] = 100
	fmt.Println("函数中修改后切片的数据", s2)
}
