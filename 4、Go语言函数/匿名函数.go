package main

import (
	"flag"
	"fmt"
)

/*
匿名函数 -》 没有函数名，只有函数体，可以被赋值给函数类型的变量
func(参数列表)(返回参数列表){
    函数体
}

*/

func main() {
	// 匿名函数的声明及赋值给变量
	f := func(data int) {
		fmt.Println("hello", data)
	}

	f(100)

	//匿名函数用作回调函数
	visit(func(i int) {
		fmt.Println(i)
	}, 1, 2, 3, 4, 7)
	//使用匿名函数实现操作封装
	/*
		flag 命令行操作
	*/
	// 使用匿名函数作为map的键值
	// 第一种写法  通过指针直接获取值
	var p string
	//命令行的声明   -skill string
	flag.StringVar(&p, "skill", "ss", "Please check you input")
	// 第二种写法 在调用时需要使用*号进行值的获取
	//ss:=flag.String("hq","","Do you need help?")
	//flag.Parse()
	/*
		以下是通过条件直接判断
	*/
	if p == "skill" {
		fmt.Println(
			"skll")
	} else if p == "fly" {
		fmt.Println("fly")
	} else {
		fmt.Println(flag.Usage)
	}
	//value:=func(data1 string){
	//	fmt.Println(data1)
	//}
	//a := *ss
	//if a =="fly"{
	//	value(a)
	//}

	/*
	   以下是通过map对输出的值进行判断
	*/
	//var skill = map[string]func(){
	//	"Fire": func() {
	//		fmt.Println("chick fire")
	//	},
	//	"run": func() {
	//		fmt.Println("soldier run")
	//	},
	//	"fly": func() {
	//		fmt.Println("angel fly")
	//	},
	//}
	//
	//if f,ok:=skill[p];ok{
	//	f()
	//}else {
	//	fmt.Println("skill not found")
	//}
	//
	//if f,ok:=skill[*ss];ok{
	//	f()
	//}else {
	//	fmt.Println("Not found")
	//}

}

func visit(f func(int), list ...int) {
	for _, v := range list {
		f(v)
	}
}
