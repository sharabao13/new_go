package main

import (
	"fmt"
)

func main() {
	//二维数组 [[10] [20]]
	multiSlice := make([][]int, 0)
	fmt.Printf("%T,%#v\n", multiSlice, multiSlice)
	multiSlice = [][]int{{10}, {10, 20}}
	fmt.Println(multiSlice)

	//三维数组 [[[3] [2]] [[1]]]
	multiSlice1 := make([][][]int, 0)
	multiSlice1 = [][][]int{{{3}, {2}}, {{1}}}
	fmt.Println(multiSlice1)

	//二维数组 [[1] [3]]
	var multiSlice2 = [][]int{{1}, {3}}
	fmt.Println(multiSlice2)

	//append 添加切片函数 [[1] [3] [10 20]]
	multiSlice2 = append(multiSlice2, []int{20, 10, 30, 40})
	fmt.Println(multiSlice2)

	multiSlice3 := []int{1, 2, 3}
	fmt.Println(multiSlice3)

	multiSlice3 = append(multiSlice3, 4, 5, 6)
	fmt.Println(multiSlice3)

	multiSlice1 = append(multiSlice1, [][]int{{1}, {2}})
	fmt.Println(multiSlice1)

}
