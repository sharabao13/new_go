package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
	//1.任务的输入
	2.任务列表
	3.任务修改
	4.任务删除
	5.详情

	任务
	id，任务名称，开始时间，结束时间，状态，负责人 map[string]string
*/
const (
	id         = "id"
	name       = "name"
	start_time = "start_time"
	end_time   = "end_time"
	status     = "status"
	user       = "user"
)

//新增任务列表
var taskInfo = make([]map[string]string, 0)

func createNewTask() {
	var input string
	taskInfoListsub := initTaskList()
	fmt.Print("任务名称: ")
	fmt.Scan(&input)
	if taskNameDeduplication(input) {
		fmt.Println("任务名重复")
		return
	} else {
		taskInfoListsub[name] = input
	}
	//fmt.Scan(&input)
	//taskInfoListsub[name] = input
	fmt.Print("开始时间: ")
	fmt.Scan(&input)
	taskInfoListsub[start_time] = input
	fmt.Print("负责人: ")
	fmt.Scan(&input)
	taskInfoListsub[user] = input
	taskInfo = append(taskInfo, taskInfoListsub)
	fmt.Println("任务添加成功")
}

//任务名去重
func taskNameDeduplication(taskname string) bool {
	for _, task := range taskInfo {
		if task[name] == taskname {
			return true
		}
	}
	return false
}

//初始化任务列表信息中的id
func initTaskListId() int {
	var taskId int
	for _, initTaskListSubId := range taskInfo {
		initTaskListSubId, _ := strconv.Atoi(initTaskListSubId["id"])
		if initTaskListSubId > taskId {
			taskId = initTaskListSubId
		}
	}
	return taskId + 1
}

//初始化任务列表信息
func initTaskList() map[string]string {
	taskInfoList := make(map[string]string)
	//id taskInfo中最大的id+1
	taskInfoList[id] = strconv.Itoa(initTaskListId())
	taskInfoList[name] = ""
	taskInfoList[start_time] = ""
	taskInfoList[end_time] = ""
	taskInfoList[status] = "新增"
	taskInfoList[user] = ""
	return taskInfoList
}

//查询任务信息
func queryTask() {
	fmt.Print("请输入查询信息: ")
	var inputText string
	fmt.Scan(&inputText)
	for _, taskKey := range taskInfo {
		if strings.Contains(taskKey[name], inputText) {
			printQueryTask(taskKey)
		}
	}
}

//打印查询信息
func printQueryTask(task map[string]string) {
	fmt.Println(strings.Repeat("-", 20))
	fmt.Println("id: ", task[id])
	fmt.Println("任务名: ", task[name])
	fmt.Println("开始时间: ", task[start_time])
	fmt.Println("结束时间: ", task[end_time])
	fmt.Println("状态: ", task[status])
	fmt.Println("用户名: ", task[user])
	fmt.Println(strings.Repeat("-", 20))
}

//声明任务信息
//var taskInfo = make([]map[string]string,0)
func main() {
	for {
		fmt.Print("请输入要操作的指令(add/query/modify/delete/exit): ")
		var inputText string
		fmt.Scan(&inputText)
		switch inputText {
		case "add":
			createNewTask()
		case "query":
			queryTask()
		case "modify":
		case "delete":
		case "exit":
			return
		default:
			fmt.Println("输入的指令有误: ")
		}
	}
}
