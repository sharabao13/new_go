package main

import (
	"fmt"
	"strconv"
)

/*
	int[][4]string{
		ip url 状态码 字节大小（B）
		{"1.1.1.1","/index.html","200","1000"},
		{"1.1.1.2","/index.html","200","10000"},
		{"1.1.1.1","/index.html","200","10000"},
	}
	1. 每个IP出现次数
	2. 每个状态码出现次数
	3. 每个IP在每个URL上产生的流量
	4.1,2,3 top10
	key,map 必须是 ==进行比较
	ip+url =>string key 可以
	[]string{ip,url} 不可以
	[2]string{ip,url} 可以
*/
func main() {
	lines := [][4]string{
		//ip url 状态码 字节大小（B）
		{"1.1.1.1", "/index.html", "200", "1000"},
		{"1.1.1.2", "/index.html", "200", "10000"},
		{"1.1.1.1", "/index.html", "200", "10000"},
	}

	ip := map[string]int{}
	status := map[int]int{}
	traffic := map[[2]string]int{}
	for _, line := range lines {
		ip[line[0]]++

		code, _ := strconv.Atoi(line[2])
		status[code]++

		key := [2]string{line[0], line[1]}
		tr, _ := strconv.Atoi(line[3])
		traffic[key] += tr

	}
	fmt.Println(ip)
	fmt.Println(status)
	fmt.Println(traffic)
	//fmt.Printf("%T\n",lines)
}
