package main

import "fmt"

func main() {
	u1 := make(map[string]string)
	u1["name"] = "xx"
	u1["age"] = "18"
	u1["name"] = "22"
	u1["age"] = "22"
	users := make(map[int]map[string]string)
	users[1] = u1
	id := getUserId1(users)
	fmt.Println(id)
}

func getUserId1(users map[int]map[string]string) int {
	var id int
	for k, _ := range users {
		if id < k {
			id = k
		}
	}
	return id
}
