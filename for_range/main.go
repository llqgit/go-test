package main

import "fmt"

type User struct {
	ID   int
	Name string
}

func getUserList() []User {
	return []User{
		{ID: 1, Name: "abc"},
		{ID: 2, Name: "123"},
	}
}

func test1() {
	userList := getUserList()
	for _, u := range userList {
		u.Name = "哈哈"
	}
	fmt.Printf("%#v\n", userList)
}

func test2() {
	userList := getUserList()
	for i := range userList {
		userList[i].Name = "哈哈"
	}
	fmt.Printf("%#v\n", userList)
}

func test3() {
	userList := getUserList()
	for i := range userList {
		u := userList[i]
		u.Name = "哈哈"
	}
	fmt.Printf("%#v\n", userList)
}

func test4() {
	userList := getUserList()
	for i := range userList {
		u := &userList[i]
		u.Name = "哈哈"
	}
	fmt.Printf("%#v\n", userList)
}

func main() {
	test1()
	test2() // 可以赋值
	test3()
	test4() // 可以赋值
}
