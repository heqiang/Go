package main

import "fmt"

//数组和切片的不同
//数组需要声明大小或者是使用... 自动计算长度
//切片不需要 [] 中没有任何符号
func main()  {
	var a  = [4] int{1,2,5}
	fmt.Println(a[1:])
	fmt.Println(a[:4])

	//声明一个int类型的切片
	var a1 [] int
	//声明一个字符串类型的切片
	var b [] string
	fmt.Print(a)
	fmt.Print(a1)
	fmt.Print(b)


	fmt.Println("下面是make内建函数的使用")
	//使用make创建一个切片 make([]type,size,cap)
	//cap 为预分配的元素数量，这个值设定后不影响 size，只是能提前分配空间，降低多次分配空间造成的性能问题。cap必须比size大
	//计算切片的大小时 还是按照size来的
	c :=make([]int,3)
	d :=make([]int,2,3)
	fmt.Println(c,d)
	fmt.Println(len(c),len(d))
}
