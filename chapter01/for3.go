package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	/*
				练习
			     1. 打印乘法口诀表

			     2. 猜数字游戏
		 			a. 生成一个0-100的随机数
		            b. 让用户最多猜5次 猜的太大提升记录剩余次数 太小提示 5次过后退出循环
	*/

	// 练习1

	//for i := 1; i <= 9; i++ {
	//	for j := 1; j <= i; j++ {
	//		fmt.Printf("%d x %d = %d\t", j, i, j*i)
	//	}
	//	fmt.Println()
	//}

	// 练习2
	//step1 生成随机数
	rand.Seed(time.Now().UnixNano())
	randNum := rand.Intn(101)
	fmt.Println(randNum)
	//step2 用户输入一个整数
	var inputNum int
	var guestTimes int
	var lastTimes int
	guestTimes = 5
	lastTimes = 1
guestStart:
	fmt.Println("请输入一个整数: ")
	fmt.Scanln(&inputNum)
	//guestStart:
	//for i := 1; i < guestTimes; i++ {
	//	if inputNum > randNum {
	//		fmt.Println(guestTimes)
	//		fmt.Printf("猜错,输入的数字比随机数大,还有%d次机会\n", guestTimes-i)
	//		guestTimes--
	//		lastTimes++
	//		fmt.Println(lastTimes)
	//		goto guestStart
	//	} else if inputNum < randNum {
	//		fmt.Printf("猜错,输入的数字比随机数5小,还有%d次机会\n", guestTimes-i)
	//		guestTimes--
	//		lastTimes++
	//		goto guestStart
	//	} else if inputNum == randNum {
	//		fmt.Println("恭喜猜中")
	//		break
	//		//goto guestStart
	//	}
	//}
	for {
		if inputNum > randNum {
			fmt.Printf("猜错,输入的数字比随机数大,还剩%d次\n", guestTimes-lastTimes)
			lastTimes++
			if lastTimes > guestTimes {
				fmt.Println("连续猜错5次...游戏结束...")
				break
			}
		} else if inputNum < randNum {
			fmt.Printf("猜错,输入的数字比随机数小,还剩%d次\n", guestTimes-lastTimes)
			lastTimes++
			if lastTimes > guestTimes {
				fmt.Println("连续猜错5次...游戏结束...")
				break
			}
		} else if inputNum == randNum {
			fmt.Println("恭喜猜中")
			break
		}
		goto guestStart
	}
}
