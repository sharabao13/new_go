package main

import "fmt"

func main() {
	//switch 其他写法
	//判断天数
	/*
			1.3 5 7 8 10 12   31填
		    4 6 9 11 30 填
		    2 闰年 29天 其他28
	*/

	var (
		mounts int
		day    int
		year   int
	)
	year = 1990
	mounts = 2
	switch mounts {
	case 1, 3, 5, 7, 8, 10, 12:
		day = 31
	case 4, 6, 9, 11:
		day = 30
	case 2:
		if year%400 == 0 || year%4 == 0 && year%100 != 0 {
			day = 29
		} else {
			day = 28
		}
	default:
		fmt.Printf("无效的月份")
	}
	fmt.Printf("%d年%d月有%d天\n", year, mounts, day)

	switch language := "golang"; language {
	case "golang":
		fmt.Println("go语言")
	case "java":
		fmt.Println("java语言")
	case "python":
		fmt.Println("python语言")
	}
}
