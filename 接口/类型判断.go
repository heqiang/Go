package main

import "fmt"

//判断一个值是不是某个类型的方法
// value ,ok := x.(T)
//x 表示一个接口的类型，T 表示一个具体的类型（也可为接口类型）。

func main() {
	var s int
	s = 12
	getType(s)
}

func getType(a interface{}) {
	switch a.(type) {
	case string:
		fmt.Println("the type is string")
	case int:
		fmt.Println("the type is int")
	case float64:
		fmt.Println("the type is float64")

	}
}
