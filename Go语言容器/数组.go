package main

import "fmt"



func main()  {
	var  a [3] int
	fmt.Println(a[0])
	for k,v := range a{
		fmt.Printf("下标是:%d,值是:%d\n",k,v)
	}
	q := [...] int {1,2,3}
	fmt.Printf("%T\n",q)
	fmt.Println(q)
	//如果两个数组类型相同，既满足一下两个条件
	// 1 数组长度  2 数组中元素类型相同的情况可以使用比较运算符（==和 !=）
	//只有当两个数组的所有元素都是相等的时候数组才是相等的，不能比较两个类型不同的数组
	/*
	多维数组
	*/
	//声明一个int类型的二维数组
	fmt.Println("下面是二维数组的程序-************")
	var array [4][2] int

	array = [4][2] int{{1,2},{4,2},{1,3},{1,5}}

	array = [4][2] int{1:{52,32},3:{52,65}}
	fmt.Print(array)

}
