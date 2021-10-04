package main

import "fmt"

func main() {
	/*
		map 与slice 结合使用
	*/
	map1 := make(map[string]string)
	map1["name"] = "王二狗"
	map1["age"] = "20"
	map1["sex"] = "男性"
	map1["address"] = "北京市xx路xx号"
	fmt.Println(map1)

	map2 := make(map[string]string)
	map2["name"] = "李小花"
	map2["age"] = "18"
	map2["sex"] = "女性"
	map2["address"] = "上海市xx路xx号"
	fmt.Println(map2)

	map3 := make(map[string]string)
	map3["name"] = "ruby"
	map3["age"] = "22"
	map3["sex"] = "女性"
	map3["address"] = "厦门xx路xx号"
	fmt.Println(map3)

	//将map放入slice中
	s1 := make([]map[string]string, 0, 3)
	s1 = append(s1, map1)
	s1 = append(s1, map2)
	s1 = append(s1, map3)
	//fmt.Println(s1)
	for k, v := range s1 {
		fmt.Printf("第%d个人的信息是: \n", k+1)
		fmt.Printf("\t姓名： %s\n", v["name"])
		fmt.Printf("\t姓名： %s\n", v["age"])
		fmt.Printf("\t姓名： %s\n", v["sex"])
		fmt.Printf("\t姓名： %s\n", v["address"])
	}
	//for k, v := range s1 {
	//	fmt.Println(k, v)
	//}
}
