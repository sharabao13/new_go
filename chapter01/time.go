package main

import (
	"fmt"
	"time"
)

func main() {
	/*
			time包
			 2000/01/02 08:10:01
		   	 2002 年
			 01 月
			 02 天
			 24进制小时 15
			 分钟 04
			 秒 05
	*/
	now := time.Now()
	fmt.Printf("%T\n", now)
	fmt.Printf("%v\n", now)
	fmt.Println("--------------")
	fmt.Println(now.Format("2000/01/02 08:10:01"))
	fmt.Println(now.Format("2000-01-02 08:10:01"))
	fmt.Println(now.Format("2000-01-02"))
	fmt.Println(now.Format("08:10:01"))

	fmt.Println(now.Unix())
	fmt.Println(now.UnixNano())
}
