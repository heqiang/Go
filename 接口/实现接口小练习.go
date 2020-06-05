package main

import "fmt"

type Animal interface {
	run()
}

type Cat struct {
	name  string
	color string
	age   int
}

type Dog struct {
	name  string
	color string
	age   int
}

func (C *Cat) run() {
	fmt.Println("猫猫是会跑的")
}
func (D *Dog) run() {
	fmt.Println("狗狗也是会跑的")
}
func main() {
	dog := &Dog{"小明", "黑色", 1}
	dog.run()
	cat := &Cat{"小白", "白色", 2}
	cat.run()

}
