package main

import (
	"fmt"
	"time"
)

type User struct {
	Name    string
	Address string
	Age     int
}

func UserInfo() User {
	return User{
		Name:    "Test",
		Address: "User Address",
		Age:     20,
	}
}

func main() {

	batch := 10000
	req := 10
	users := make([]User, 0, batch)

	for i := 0; i < batch; i++ {
		users = append(users, UserInfo())
	}

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for i := 0; i < len(users); i += req {
		<-ticker.C

		end := i + req
		if end > len(users) {
			end = len(users)
		}
    
		for j := i; j < end; j++ {
			fmt.Println(users[j])
		}
    
	}

}
