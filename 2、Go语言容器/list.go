package main

import (
	"container/list"
	"fmt"
)

/*
list初始化的两种方式
1  Name := list.New()
2  var Name list.List
*/

func main() {
	l := list.List{}
	//插到队列的后方
	l.PushBack("First")
	//插到队列的前方
	l.PushFront(122)

	element := l.PushBack("ele")
	fmt.Println(element.Value)
	l.InsertAfter("After", element)
	l.InsertBefore("Before", element)
	//list遍历
	for i := l.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
}
