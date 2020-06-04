package main

import "fmt"

func fact(num int) (res int) {
	if num > 0 {
		result := num * fact(num-1)
		return result
	}
	return 1
}

func main() {
	res := fact(5)
	fmt.Println(res)
}
