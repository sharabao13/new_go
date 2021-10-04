package main

import (
	"fmt"
)

func main() {
	/*
			slice 切片
			切片是数组的抽象
			go数组中长度不可变，在特定场景中这样的集合不太适用，go提供了一种灵活，功能强悍的内置类型切片（动态数组）
			与数组比切片的长度是不固定的，可以追加元素，在追加事可能使切片的容量增加
			切片是一种方便的、灵活且强大的包装器，切片本身没有任何数据。他们只是对现有数组的引用
			切片与数组比，不需要设定长度，在[]中不用设定值，相对来说比较自由
			从概念上面来说slice像一个结构体,这个结构体包含三个元素
			1.指针 指向数组中slice指定的开始位置
			2.长度 slice的长度
			3.最大长度 slice开始位置到数组的最后位置的长度

			数组array
				存储一组相同数据类型的数据结构
			切片
				同数组类似，也叫变长数组或者动态数组
				特点 变长
			    是一个引用类型的容器，指向了一个底层的数组

			语法
			var indetieier []type  不需要写长度

			常用make函数来创建数据
			func make(t Type, size ...IntegerType) Type
			 第一个参数 类型 可以是slice map chann
		     第二个参数 长度 实际存储元素的数量
		 	 第三个参数 容量 最多能存储元素的数量


	*/
	// 数组
	arr1 := [4]int{1, 2, 3, 4}
	fmt.Println(arr1)

	// 切片
	var s1 []int
	fmt.Println(s1)

	s2 := []int{1, 2, 3, 4} //边长
	fmt.Println(s2)
	fmt.Printf("数组的类型%T,切片的类型%T\n", arr1, s2)
	fmt.Println(len(s2))
	fmt.Println(cap(s2))

	//make创建slice
	s3 := make([]int, 2, 4)
	fmt.Println(s3)
	fmt.Printf("长度%d,容量%d\n", len(s3), cap(s3))

	//slice赋值
	s3[0] = 1
	s3[1] = 2
	//s3[2] = 3 panic: runtime error: index out of range [2] with length 2
	fmt.Println(s3)

	//append 函数 slice末尾追加元素
	//func append(slice []Type, elems ...Type) []Type
	//slice = append(slice, elem1, elem2)
	//slice = append(slice, anotherSlice...)
	s4 := make([]int, 0, 4)
	s4 = append(s4, 1, 2)
	fmt.Println(s4)
	s4 = append(s4, s3...)
	fmt.Println(s4)
	fmt.Println("--------------")

	//slice 遍历
	for i := 0; i < len(s4); i++ {
		fmt.Println(s4[i])
	}

	//for range
	for i1, v1 := range s4 {
		fmt.Printf("索引%d,元素值%d\n", i1, v1)
	}

}
