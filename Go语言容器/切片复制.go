package main

import "fmt"

func main() {
	//var  src  = []int{1,5,3,6}
	src1 := make([]int, 20)
	dest := make([]int, 20)

	for i := 0; i < 15; i++ {
		src1[i] = i
	}
	fmt.Println(src1)
	copy(dest, src1)
	fmt.Println(dest)
}
