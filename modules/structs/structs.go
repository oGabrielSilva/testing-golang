package main

import (
	"fmt"
	"time"
)

type User struct {
	Name      string
	Age       int
	Email     string
	CreatedAt time.Time
}

func main() {
	user := User{Name: "Gabriel", Age: 23, Email: "gabriel04gh1.gh@gmail.com", CreatedAt: time.Now().Add(time.Hour)}
	user2 := User{Email: "gabriel04gh1.gh@gmail.com", Name: "Gabriel", CreatedAt: time.Now().Add(time.Hour)}

	fmt.Println("Usuário", user.Email)
	fmt.Println(user)
	fmt.Println(user2)
}
