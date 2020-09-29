package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age  int
}

type UserSlice []User

func (userSlice UserSlice) Len() int {
	return len(userSlice)
}

func (userSlice UserSlice) Less(i int, j int) bool {
	return userSlice[i].Age < userSlice[j].Age
}

func (userSlice UserSlice) Swap(i, j int) {
	userSlice[i], userSlice[j] = userSlice[j], userSlice[i]
}

func main() {
	users := []User{
		{"Hury", 23},
		{"Budi", 45},
		{"Toni", 10},
		{"Hulk", 30},
		{"IronMan", 60},
	}
	fmt.Println("Sebelum di sort:", users)
	sort.Sort(UserSlice(users))
	fmt.Println("Sesudah di sort:", users)
}
