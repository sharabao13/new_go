package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	/*
		用户访问界面
		输入密码 3次错误 退出
		用户管理系统
		命令行的用户管理
		用户信息存储
			内存
			结构  slice map
			用户 id ,name age tel addr
					 map  暂用string
		用户添加
		用户修改
			// 请输入需要修改的用户ID
			// user[id] => 判断存在一已否
			//打印用户信息,提示用户是否修改
			// 用户输入修改内容
		用户删除
		用户的查询
	*/
	//存储用户信息
	loginPage()
	users := make(map[string]map[string]string)
	userId := 0
	//fmt.Println("欢迎使用XXX用户系统")

	for {
		var op string
		fmt.Println(`
1. 新建用户
2. 修改用户
3. 删除用户
4. 查询用户
5. 退出
请输入您的指令: `)
		fmt.Scanln(&op)
		//fmt.Println("请输入您的指令: ", op)
		switch op {
		case "1":
			userId++
			addUser(userId, users)
		case "2":
			changeUser(users)
		case "3":
			deleteUsers(users)
		case "4":
			selectUser(users)
		case "5":
			fmt.Println("退出系统")
			return
		default:
			fmt.Println("输入有误")
		}

	}

}

//添加用户
func addUser(userId int, users map[string]map[string]string) {
	var (
		id   string
		name string
		age  string
		tel  string
		addr string
	)
	id = strconv.Itoa(userId)
	//userId = strconv.Atoi(id)
	//fmt.Println(id)
	fmt.Println("请输入您的姓名: ")
	fmt.Scanln(&name)
	fmt.Println("请输入您的年龄: ")
	fmt.Scanln(&age)
	fmt.Println("请输入您的电话: ")
	fmt.Scanln(&tel)
	fmt.Println("请输入您的住址: ")
	fmt.Scanln(&addr)
	users[id] = map[string]string{
		"id":   id,
		"name": name,
		"age":  age,
		"tel":  tel,
		"addr": addr,
	}
	fmt.Println(users)
}

//查询用户
func selectUser(users map[string]map[string]string) {
	//queryInfo 为空查全部
	//非空，name,age,tel addr 任一一个对象中包含queryInfo内容的显示
	var queryInfo string
	fmt.Println("请输入查询信息: ")
	fmt.Scanln(&queryInfo)
	titel := fmt.Sprintf("%5s|%10s|%3s|%10s|%20s", "Id", "Name", "AGE", "Tel", "Addr")
	fmt.Println(titel)
	fmt.Println(strings.Repeat("-", len(titel)))
	for _, user := range users {
		if queryInfo == "" || strings.Contains(user["name"], queryInfo) || strings.Contains(user["age"], queryInfo) || strings.Contains(user["tel"], queryInfo) || strings.Contains(user["addr"], queryInfo) {
			fmt.Printf("%5s|%10s|%3s|%10s|%20s", user["id"], user["name"], user["age"], user["tel"], user["addr"])
		}
	}
}

//修改用户
func changeUser(users map[string]map[string]string) {
	var userId string
	fmt.Println("请输入要修改的用户ID： ")
	fmt.Scanln(&userId)
	if user, ok := users[userId]; ok {
		fmt.Println("您将修改的用户信息是: ")
		fmt.Println(user)
		var sureInfo string
		fmt.Println("确定修改Y/N")
		fmt.Scanln(&sureInfo)
		if sureInfo == "y" || sureInfo == "Y" {
			var (
				name string
				age  string
				tel  string
				addr string
			)
			fmt.Println("请输入您的姓名: ")
			fmt.Scanln(&name)
			fmt.Println("请输入您的年龄: ")
			fmt.Scanln(&age)
			fmt.Println("请输入您的电话: ")
			fmt.Scanln(&tel)
			fmt.Println("请输入您的住址: ")
			fmt.Scanln(&addr)
			users[userId] = map[string]string{
				"name": name,
				"age":  age,
				"tel":  tel,
				"addr": addr,
			}

		}
	}
}

//删除用户
func deleteUsers(users map[string]map[string]string) {
	var userId string
	fmt.Println("请输入要删除的用户ID： ")
	fmt.Scanln(&userId)
	if user, ok := users[userId]; ok {
		fmt.Println("您将修改的用户信息是: ")
		fmt.Println(user)
		var sureInfo string
		fmt.Println("确定修改Y/N")
		fmt.Scanln(&sureInfo)
		if sureInfo == "y" || sureInfo == "Y" {
			delete(users, userId)
		}
	} else {
		fmt.Println("Id 不存在")
	}
}

//登录界面
func loginPage() {
	fmt.Println("欢迎使用xxx用户管理系统")
	const passWord string = "Admin@123"
	const inputTimes int = 3
	var userPassword string
	for i := 0; i < inputTimes; i++ {
		fmt.Println("请输入密码: ")
		fmt.Scanln(&userPassword)
		if userPassword == passWord {
			fmt.Println("登录成功")
			break
		} else {
			fmt.Println("密码错误")
		}
		if i == inputTimes-1 {
			fmt.Println("连续3次错误,退出")
			return
		}
	}

}
