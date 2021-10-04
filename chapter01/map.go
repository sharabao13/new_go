package main

import "fmt"

func main() {
	/*
			map 集合
				map是go中的内置类型，它将一个值与一个键关联起来。可以使用相应的键索引值
				map是一种无序的键值对的集合。map最重要的一点是通过key来快速检索数据，key类似索引，指向数据的值
				map是一种集合，所以我们可以像迭代数组和切片去那样迭代它，不过map是无序的，我们无法决定它的返回顺序，这是因为
				map是使用hash表来实现的，也是引用类型
				注：
					map是无序的，每次打印出来的map都会不一样，它不能通过index获取，而必须通过key获取
					map的长度是不固定的，也就是和slice一样，也是一种引用类型
					内置的len函数同样适用map，返回map拥有key的数量
					map的key可以是所有可比较的类型，bool int float ...

				语法结构
				1. 创建map
					var map1 map[key类型]value类型
					  nil map 无法直接使用
					var map2 = make(map[key类型]value类型)
		         	var map3 = map[key类型]value类型{key:value,key:value}

				2. 添加修改
					map[key] = value
					如果key存在，就是添加数据
					如果key不存在，就是修改数据

				3. 获取
					map[key]  --> value
					value,ok := map[key]
						根据key获取对应的value
							如果key存在，value就是对应的数据,ok 为true
							如果key不存在, value就是值类型的默认值，ok为false

				4. 删除数据
					delete函数
					delete(map,key)
						如果key存在，直接删除
						如果不存在，删除失败

				5. 长度
					len（map）
	*/

	var map1 map[int]string
	var map2 = make(map[int]string)
	var map3 = map[string]int{"go": 98, "python": 78, "java": 80}
	fmt.Println(map1)
	fmt.Println(map2)
	fmt.Println(map3)

	fmt.Println(map1 == nil)
	fmt.Println(map2 == nil)
	fmt.Println(map3 == nil)
	//map1[1] = "hello"  //nil map不能直接使用
	if map1 == nil {
		map1 = make(map[int]string)
	}

	//添加
	map1[1] = "hello"
	map1[2] = "world"
	map1[3] = "hcx"
	map1[4] = "gfb"
	map1[5] = "tae"
	map1[6] = "ruby"
	fmt.Println(map1)

	//获取
	fmt.Println(map1[2])
	fmt.Println(map1[20])

	v1, ok := map1[2]
	if ok {
		fmt.Println("对应的值是:", v1)
	} else {
		fmt.Println("对应的key不存在值是:", v1)
	}

	//修改
	map1[5] = "Tae"
	fmt.Println(map1)

	//删除key
	delete(map1, 6)
	fmt.Println(map1)

	//长度
	fmt.Println(len(map1))

	map4 := map[int]string{1: "ss", 2: "aa"}
	if v1, ok := map4[20]; ok {
		fmt.Println("存在", v1)
	} else {
		fmt.Println("不存在", v1)
	}
}
