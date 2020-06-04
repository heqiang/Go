package main

import (
	"encoding/json"
	"fmt"
)

//屏幕
type Screen struct {
	Size int
}

type Battery struct {
	Capacity int
}

func getJsonData() []byte {
	raw := &struct {
		Screen
		Battery
		HasTochID bool
	}{
		Screen{
			12,
		},
		Battery{
			2910,
		},
		true,
	}
	//json.Marshal() json序列化 将raw变量序列化为[]byte 类型的json数据
	jsonData, _ := json.Marshal(raw)
	return jsonData
}

func main() {
	jsonData := getJsonData()
	fmt.Println(string(jsonData))
}
