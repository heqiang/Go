package main

import "fmt"

func main() {
	//var i int
	//for i=1;i<10;i++{
	//	for j:=1;j<=i;j++{
	//		fmt.Printf("%d*%d=%d\t",j,i,i*j)
	//	}
	//	fmt.Println()
	//}
	for x := 0; x < 10; x++ {
		for y := 0; y < 10; y++ {
			if y == 2 {
				goto breakHere
			}
		}
	breakHere:
		fmt.Print("helo")
	}

}
