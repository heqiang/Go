package main

import "fmt"

func swap(a ,b *int){
		temp := *a
		*a = *b
		*b = temp
}


func  main()  {
	a,b := 2,3
	swap(&a,&b)
	fmt.Println(a)
	fmt.Print(b)

	str := new(string)
	*str = "Go语言教程"
	fmt.Println(*str)
}