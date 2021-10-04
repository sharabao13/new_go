package main

import "fmt"

func main() {
	/*
			数组的遍历
		 		方法一  依次访问数组中的元素
					arr[0],arr[1],arr[2],arr[3]
		        方法二  通过循环，配合下标
					for i:=0;i<len(arr);i++{
							arr[i]
			 		}
		 		方法三  使用range
					range, 定义范围
					不需要操作数组的下标，到达数组的末尾，自动结束 for range 循环
					  每次都数组鸿获取下标和对应的数值
					  如果只需要值并希望忽略索引，那么可以通过使用_ blank标识符替换索引

	*/
	// 方法1
	arr1 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(arr1[0])
	fmt.Println(arr1[1])
	fmt.Println(arr1[2])
	fmt.Println(arr1[3])
	fmt.Println(arr1[4])
	fmt.Println("------------------")

	// 方法二
	for i := 0; i < len(arr1); i++ {
		fmt.Println(arr1[i])
	}
	fmt.Println("------------------")

	//方法三
	for index, value := range arr1 {
		fmt.Printf("数组的索引为: %d,值为: %d\n", index, value)
	}
	fmt.Println("------------------")
	//不显示索引
	for _, value := range arr1 {
		fmt.Printf("数组的值为: %d\n", value)
	}
	fmt.Println("------------------")
}
