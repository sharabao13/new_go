package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//猜数字游戏
	//step1 生成随机数
	rand.Seed(time.Now().UnixNano())
	randNum := rand.Intn(100)
	fmt.Println(randNum)
	//setp2 for循环
	const maxGuestTImes = 5
	//for i := 0; i < maxGuestTImes; i++ {
	//	var inputNum int
	//	fmt.Println("请输入一个整数")
	//	fmt.Scanln(&inputNum)
	//	if inputNum == randNum {
	//		fmt.Printf("恭喜猜中,用了%d次机会", i+1)
	//		return
	//	} else if inputNum > randNum {
	//		fmt.Printf("猜的数字比随机数大,还剩%d机会\n", maxGuestTImes-i-1)
	//	} else if inputNum < randNum {
	//		fmt.Printf("猜的数字比随机数小,还剩%d机会\n", maxGuestTImes-i-1)
	//	}
	//	if i == maxGuestTImes-1 {
	//		fmt.Printf("%d次机会没猜中,游戏结束", maxGuestTImes)
	//	}
	//}
	for {
		for i := 0; i < maxGuestTImes; i++ {
			var inputNum int
			fmt.Println("请输入一个整数")
			fmt.Scanln(&inputNum)
			if inputNum == randNum {
				fmt.Printf("恭喜猜中,用了%d次机会", i+1)
				return
			} else if inputNum > randNum {
				fmt.Printf("猜的数字比随机数大,还剩%d机会\n", maxGuestTImes-i-1)
			} else if inputNum < randNum {
				fmt.Printf("猜的数字比随机数小,还剩%d机会\n", maxGuestTImes-i-1)
			}
			if i == maxGuestTImes-1 {
				fmt.Printf("%d次机会没猜中,游戏结束\n", maxGuestTImes)
			}
		}
		var secondGust string
		fmt.Println("重新开始吗: (y/n)")
		fmt.Scanln(&secondGust)
		if secondGust != "y" {
			return
		}
	}
}
