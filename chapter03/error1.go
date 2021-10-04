package main

import (
	"fmt"
	"net"
	"path/filepath"
)

func main() {
	addr, err := net.LookupHost("www.sina.com")
	fmt.Println(err)
	if ins, ok := err.(*net.DNSError); ok {
		if ins.Timeout() {
			fmt.Println("超时")
		} else if ins.Temporary() {
			fmt.Println("临时性错误")
		} else {
			fmt.Println("一般性错误")
		}
	}
	fmt.Println(addr)

	files, err := filepath.Glob("")
	if err != nil && err == filepath.ErrBadPattern {
		fmt.Println(err)
		return
	}
	fmt.Println(files)
}
