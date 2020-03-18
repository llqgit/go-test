package main

import "fmt"

type User struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func main() {

	userMap := make(map[int64]User)

	userMap[0] = User{
		ID:   1,
		Name: "1",
	}
	userMap[1] = User{
		ID:   1,
		Name: "1",
	}
	userMap[2] = User{
		ID:   1,
		Name: "1",
	}

	fmt.Println(userMap)

	// === 这种写法是错误的 ===
	// go 修改 map 元素内容需要整体赋值，不能单独修改元素的子内容
	for i := int64(0); i < 10; i++ {
		if user, ok := userMap[0]; ok {
			user.ID += 1
		}
	}

	fmt.Println("错误map", userMap)

	// +++ 正确的写法 +++
	for i := int64(0); i < 10; i++ {
		if user, ok := userMap[0]; ok {
			userMap[0] = User{
				ID:   user.ID + 1,
				Name: user.Name,
			}
		}
	}

	fmt.Println("正确map", userMap)
}
