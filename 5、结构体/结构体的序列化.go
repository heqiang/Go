package main

import (
	"encoding/json"
	"fmt"
)

/*
json序列化是指，将有 key-value 结构的数据类型（比如结构体，map，切片）序列化成json字符串的操作
*/

type Hero struct {
	name     string `json:"hq_name"`
	Age      int    `json:"hq_agq"`
	Birthday string
	Sal      float64
	Skill    string
}

func teststruct() {
	//结构体的序列化
	hero := Hero{
		name:     "hq",
		Age:      18,
		Birthday: "2009-11-11",
		Sal:      8000.0,
		Skill:    "Coding",
	}
	//将
	data, err := json.Marshal(hero)
	if err != nil {
		fmt.Printf("序列号错误 err=%v\n", err)
	}
	fmt.Printf("结构体序列化后=%v\n", string(data))

	//map的序列化
	a := make(map[string]interface{})
	a["name"] = "hq"
	a["age"] = 14
	a["address"] = "四川"
	jsonData2, err := json.Marshal(a)
	if err != nil {
		fmt.Printf("map序列化错误 err=%v\n", err)
	}
	fmt.Printf("map 序列化结果 %v\n", string(jsonData2))
	//切片数据序列化

	sliceJson := make(map[string]interface{})
	fmt.Println(sliceJson)

}

func main() {
	teststruct()
}
