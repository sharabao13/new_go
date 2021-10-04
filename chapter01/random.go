package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	/*
				 	 生成随机数 random
					  伪随机数，根据一定的算法公式算出来
				 	  math/rand
			          rand.sead 生成种子数

		               step1: 设置种子数
		               step2: 调用生成随机数的函数
	*/

	num1 := rand.Int()
	fmt.Println(num1)

	for i := 1; i < 5; i++ {
		num := rand.Intn(10)
		fmt.Println(num)
	}

	rand.Seed(time.Now().UnixNano())
	for i := 1; i <= 5; i++ {
		fmt.Println(rand.Intn(11))
	}
}
