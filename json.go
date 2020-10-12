package main

import (
	"encoding/json"
	"fmt"
)

type UserJson struct {
	Fullname string `json:"name"`
	Age      int
}

func main() {
	var jsonString = `{"Name": "Mashuri Mansur", "Age": 24}`
	var jsonData = []byte(jsonString)

	var data UserJson

	var err = json.Unmarshal(jsonData, &data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("User :", data.Fullname)
	fmt.Println("Age :", data.Age)

	var data1 map[string]interface{}
	_ = json.Unmarshal(jsonData, &data1)
	fmt.Println("User :", data1["Name"])
	fmt.Println("Age :", data1["Age"])

	var jsonStringArray = `[
		{"Name": "john wick", "Age": 27},
		{"Name": "ethan hunt", "Age": 32}
	]`

	var data3 []UserJson
	json.Unmarshal([]byte(jsonStringArray), &data3)
	fmt.Println("user 1:", data3[0].Fullname)
	fmt.Println("user 2:", data3[1].Fullname)

	var object = []UserJson{{"john wick", 27}, {"ethan hunt", 32}}
	var jsonData4, _ = json.Marshal(object)
	var jsonString4 = string(jsonData4)
	fmt.Println(jsonString4)
}
