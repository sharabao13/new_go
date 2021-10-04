package main

import (
	"fmt"
	"strings"
)

func main() {
	/*
		go中的字符串是一个字节切片，可以通过将其内容封装在""中创建字符串,
		strings 包的使用
	*/
	s1 := "helloworld"
	//contains 包含指定的内容 返回bool
	fmt.Println(strings.Contains(s1, "ella"))
	//containsany 包含指定的内容的任一字符即可 返回bool
	fmt.Println(strings.ContainsAny(s1, "ella"))
	//conuts 统计包含内容出现的次数 返回int
	fmt.Println(strings.Count(s1, "l"))

	//hasprefix 开头 返回bool
	s2 := "202105课堂笔记.txt"
	if strings.HasPrefix(s2, "2021") {
		fmt.Println("2021年开头的文件")
	}
	//hassuffix结尾 返回bool
	if strings.HasSuffix(s2, "txt") {
		fmt.Println("txt文件的结尾")
	}

	fmt.Println(strings.Index(s1, "llo"))   //查找substr在s中的位置，如果不存在返回-1
	fmt.Println(strings.IndexAny(s1, "dh")) //查找chars中任一字符，出现在s中的位置

	//字符串拼接  join 切片操作
	ss1 := []string{"abc", "hello", "world"}
	s3 := strings.Join(ss1, "-")
	fmt.Println(s3)

	//切割 返回一个切片
	s4 := "123,222,sadd,aaa,bbb,12d"
	ss4 := strings.Split(s4, ",")
	for _, v := range ss4 {
		fmt.Println(v)
	}

	//重复 自己拼接自己count次
	s5 := strings.Repeat("hello", 3)
	fmt.Println(s5)

	// 替换 n 替换的个数，-1 替换所有
	s6 := strings.Replace(s1, "l", "e", -1)
	fmt.Println(s6)

	//大小写
	s7 := strings.ToUpper(s1) // 大写
	s8 := strings.ToLower(s1) //小写
	fmt.Println(s7, s8)
}
