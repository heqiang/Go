package main

import "fmt"

func  main()  {
	 	str := "welcome to Go World"+"hq"
	 	str1 := `
			第一行
			第二行
			第三行
			`
		str3 := &str1
		str4 := *str3
		fmt.Print(str)
	 	fmt.Print(str1)
		fmt.Print(str4)
}
