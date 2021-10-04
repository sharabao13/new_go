package main

import "fmt"

func main() {
	/*
		深拷贝： 拷贝的是数据本身
			值类型的数据，默认都是深拷贝 array int float string bool

		浅拷贝： 拷贝的是内存地址
			引用类型的数据默认都是浅拷贝 slice map..
	*/

	// for 循环 切片深拷贝
	s1 := []int{1, 2, 3, 4}
	s2 := make([]int, 0)
	for i := 0; i < len(s1); i++ {
		s2 = append(s2, s1[i])
	}
	fmt.Println(s1)
	fmt.Println(s1)

	//copy 函数
	//func copy(dst, src []Type) int
	s3 := []int{7, 8, 9}
	fmt.Println(s2)
	fmt.Println(s3)

	//copy(s2, s3) //将s3中的元素，拷贝到s2中 从s2的第一个元素开始覆盖
	//copy(s3, s2)   //将s2中的元素，拷贝到s3中 从s3的第一个元素开始覆盖，容量不过不复制剩余的元素
	copy(s2[2:], s3)
	fmt.Println(s2)
	fmt.Println(s3)
}
