package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	maxAuth       = 3
	loginPassword = "Admin@123"
)

//获取ID
func getUserId(users map[int]map[string]string) int {
	var id int
	for k, _ := range users {
		if id < k {
			id = k
		}
	}
	return id + 1
}

// 登录验证
func loginAuth() bool {
	var inputPass string
	for i := 1; i <= maxAuth; i++ {
		fmt.Println("请输入密码: ")
		fmt.Scanln(&inputPass)
		if inputPass == loginPassword {
			return true
		} else {
			fmt.Println("密码错误")
		}
	}
	return false
}

//添加用户信息
func userAddInfo() map[string]string {
	user := map[string]string{}
	user["name"] = inputInfo("请输入您的姓名: ")
	user["age"] = inputInfo("请输入您的年龄: ")
	user["tel"] = inputInfo("请输入您的电话: ")
	user["addr"] = inputInfo("请输入您的住址: ")
	return user
}

//用户添加
func userAdd(users map[int]map[string]string) {
	id := getUserId(users)
	user := userAddInfo()
	users[id] = user
	fmt.Println("添加成功")
}

//删除用户
func deleteUser(users map[int]map[string]string) {
	if id, err := strconv.Atoi(inputInfo("请输入要删除ID: ")); err == nil {
		if user, ok := users[id]; ok {
			fmt.Println("您将删除的用户是: ")
			fmt.Println(user)
			sureInfo := inputInfo("是否确定删除Y/N")
			if sureInfo == "y" || sureInfo == "Y" {
				delete(users, id)
				fmt.Println("删除成功")
			}
		} else {
			fmt.Println("输入的ID不正确")
		}
	} else {
		fmt.Println("输入的ID不正确")
	}
}

//修改用户
func userChange(users map[int]map[string]string) {
	if id, err := strconv.Atoi(inputInfo("请输入要修改ID: ")); err == nil {
		if user, ok := users[id]; ok {
			fmt.Println("您将修改的用户是: ")
			fmt.Println(user)
			sureInfo := inputInfo("是否确定修改Y/N")
			if sureInfo == "y" || sureInfo == "Y" {
				user := userAddInfo()
				users[id] = user
				fmt.Println("修改成功")
			}
		} else {
			fmt.Println("输入的ID不正确")
		}
	} else {
		fmt.Println("输入的ID不正确")
	}
}

//输入信息
func inputInfo(userInput string) string {
	var input string
	fmt.Println(userInput)
	fmt.Scanln(&input)
	return strings.TrimSpace(input)
}

//用户查询
func userQuery(users map[int]map[string]string) {
	query := inputInfo("请输入要查询的信息")
	titel := fmt.Sprintf("%3s|%10s|%3s|%10s|%20s", "Id", "Name", "AGE", "Tel", "Addr")
	fmt.Println(titel)
	fmt.Println(strings.Repeat("-", len(titel)))
	for id, user := range users {
		if query == "" || strings.Contains(user["name"], query) || strings.Contains(user["age"], query) || strings.Contains(user["tel"], query) || strings.Contains(user["addr"], query) {
			fmt.Printf("%3d|%10s|%3s|%10s|%20s", id, user["name"], user["age"], user["tel"], user["addr"])
			//fmt.Println(id, user)
		}
	}
}
func main() {
	if !loginAuth() {
		fmt.Printf("密码%d次错误，程序退出\n", maxAuth)
		return
	}
	menu := `
*********************************
  1. 新建用户
  2. 修改用户
  3. 删除用户
  4. 查询用户
  5. 退出
*********************************
`
	users := make(map[int]map[string]string)
	callBacks := map[string]func(map[int]map[string]string){
		"1": userAdd,
		"2": userChange,
		"3": deleteUser,
		"4": userQuery,
		"5": func(m map[int]map[string]string) {
			os.Exit(0)
		},
	}
	fmt.Println("欢迎进入xxx用户管理系统")
	//END:
	for {
		fmt.Println(menu)
		oper := inputInfo("请输入指令")
		callback, ok := callBacks[oper]
		if ok {
			callback(users)
		} else {
			fmt.Println("指令错误")
		}
		//switch oper {
		//case "1":
		//	userAdd(users)
		//case "2":
		//	userChange(users)
		//case "3":
		//	deleteUser(users)
		//case "4":
		//	userQuery(users)
		//case "5":
		//	break END
		//default:
		//	fmt.Println("指令错误")
		//
		//}
	}
}
