package main

import (
	"fmt"
	"sort"
)

func main() {
	/*
		map 遍历
			for range
	*/
	map1 := make(map[int]string)
	map1[1] = "红孩儿"
	map1[2] = "小转风"
	map1[3] = "白骨精"
	map1[4] = "白素贞"
	map1[5] = "金角大王"
	map1[6] = "王二狗"

	for k, v := range map1 {
		fmt.Println(k, v) //无序的
	}
	fmt.Println("------------------------")
	// 有序的打印map key是int类型

	//step1 获取所有的key添加到 slice array
	keys := make([]int, 0, len(map1))
	for k, _ := range map1 {
		keys = append(keys, k)
	}
	fmt.Println(keys)

	//step2 sort 排序 func Ints(x []int)
	sort.Ints(keys)
	fmt.Println(keys)

	//遍历key ——————》map[key]
	for _, k := range keys {
		fmt.Println(k, map1[k])
	}
}
