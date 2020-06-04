package main

import "fmt"

func fib(num int) (res int) {
	if num < 2 {
		res = 1
	} else {
		res = fib(num-1) + fib(num-2)
	}
	return res
}

func main() {

	for x := 0; x < 10; x++ {
		result := fib(x)
		fmt.Printf("fib(%d) is %d\n", x, result)
	}
}
