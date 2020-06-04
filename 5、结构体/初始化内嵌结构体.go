package main

import "fmt"

//车轮
type Wheel struct {
	Size int
}

//引擎
type Engine struct {
	Power int
	Type  string
}

//车
type Car struct {
	Wheel
	Engine
}

func main() {
	c := Car{
		Wheel{
			Size: 18,
		},
		Engine{
			Type:  "1.7t",
			Power: 142,
		},
	}
	fmt.Println(c.Size)
}
