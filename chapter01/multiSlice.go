package main

import (
	"fmt"
)

func main() {
	multiSlice := make([][]int, 0)
	fmt.Printf("%T,%#v\n", multiSlice, multiSlice)
	multiSlice = [][]int{{10}, {10, 20}}
	fmt.Println(multiSlice)

	multiSlice1 := make([][][]int, 0)
	multiSlice1 = [][][]int{{{3}, {2}}, {{1}}}
	fmt.Println(multiSlice1)
}
