package main

import "fmt"

/*
函数名、形式参数列表、返回值列表（可省略）以及函数体
func 函数名(形式参数列表)(返回值列表){
    函数体
}
*/

func Get_num(a ...int) {
	fmt.Println(a)
	for _, v := range a {
		fmt.Println(v)
	}
}

func main() {
	Get_num(1, 2, 3, 4, 5)

}
