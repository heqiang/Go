package main

import "fmt"

func Accu(value int) func() int {
	return func() int {
		value++

		return value
	}
}

//创建一个玩家生成器，输入名称，输出生成器
func playGame(name string) func() (string, int) {
	HP := 50
	return func() (string, int) {
		return name, HP
	}
}

func main() {
	accumulator := Accu(1)

	fmt.Println(accumulator())
	fmt.Println(accumulator())
	fmt.Println(accumulator())

	//创建一个玩家生成器
	generator := playGame("hq")
	name, hp := generator()
	fmt.Println(name, hp)

}
