package main

import "fmt"

func main() {
	/*
		数组指针 首先是一个指针，一个数组的地址
			*[4]type

		指针数组 首先是一个数组，存储的数据类型是指针
			[4]*type

		*[5]float64  指针，一个存储了float64类型的数组指针
		*[3]string	 指针，一个存储了string类型的数组指针
		[3]*string   数组，一个存储了string类型的指针数组
		[5]*float64  数组，一个存储了float类型的指针数组
		*[5]*float54 指针，一个存储了float类型的指针地址的数组指针
		*[3]*string  指针，一个存储了string类型的指针地址的数组指针
		**[4]string  指针，存储了string类型的指针的指针
		**[4]*string 指针，存储了string类型指针地址的指针的指针
	*/
	//普通数组
	arr1 := [4]int{1, 2, 3, 4}
	fmt.Println(arr1)

	//创建一个指针，存储的是数组的地址，  数组指针
	var p1 *[4]int
	p1 = &arr1
	fmt.Println(p1)
	fmt.Printf("%p\n", p1)  //数组的地址
	fmt.Printf("%p\n", &p1) //p1指针的地址
	fmt.Printf("%v\n", *p1) //p1指针的地址对应数组的值

	//根据数组指针，操作数组
	(*p1)[0] = 100
	fmt.Println(arr1)
	//简写
	p1[0] = 200
	fmt.Println(arr1)

	//指针数组
	a := 1
	b := 2
	c := 3
	d := 4
	arr2 := [4]int{a, b, c, d}
	arr3 := [4]*int{&a, &b, &c, &d}
	fmt.Println(arr2)
	fmt.Println(arr3)

	arr2[0] = 100
	fmt.Println(arr2)
	fmt.Println(a)

	*arr3[0] = 200
	fmt.Println(arr3)
	fmt.Println(a)

	for i := 0; i < len(arr3); i++ {
		fmt.Println(*arr3[i])
	}
}
