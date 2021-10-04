package main

import "fmt"

func main() {
	/*
			数组
		    定义： 具有相同唯一类型的一组已编号且长度固定的数据项序列，这种类型可以是任意的原始类型，例如整形、字符串、布尔或者自定义类型
		    数组元素可以通过索引[位置] 来读取 或者修改，索引从0开始，第一个元素索引为0，第二为1，以此类推，
		    数组的下标范围是从0开始导长度-1
		    数组一旦定义了，大小不能更改

		    语法
		    声明和初始化数组
		    需要指明数组的大小和存储的数据类型
		    var variable_name [SIZE] variable_type
		    var blance [10] float32
		    var blance = [5] float32{1000.0,2.0.3.0.7.0.50.0}
	*/
	//step1 创建数组
	var arr1 [4]int
	arr1[0] = 1
	arr1[1] = 2
	arr1[2] = 3
	arr1[3] = 4
	fmt.Println(arr1[0]) //打印第一个数组
	fmt.Println(arr1[1]) //打印第二个数组

	fmt.Println(len(arr1)) //容器中实际存储的数据量
	fmt.Println(cap(arr1)) //容器中能够存储的最大数量

	//修改数组数据 通过下标修改
	arr1[1] = 100
	fmt.Println(arr1)

	//其他数据声明跟初始化方式 为赋值的数组默认为空值
	var a [4]int
	fmt.Println(a)

	var b = [4]int{1, 2, 3, 4}
	fmt.Println(b)

	var c = [4]int{1: 1, 2: 2}
	fmt.Println(c)

	var e = [3]string{"tae", "bao", "hcx"}
	fmt.Println(e)

	//未写元素长度默认为赋值的长度
	f := [...]int{1, 2, 3, 4, 5}
	fmt.Println(f)

}
