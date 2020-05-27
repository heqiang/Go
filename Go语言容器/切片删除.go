package main

import "fmt"

func main() {

	slice := []int{1, 3, 5}
	//使用opend删除开头的N个元素
	var N int
	slice = append(slice[:0], slice[N:]...)

	//使用copy删除开头的N个元素
	slice = slice[:3]

	fmt.Print(slice)
}
