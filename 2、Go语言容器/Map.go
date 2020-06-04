package main

import "fmt"

//map
// var mapName map[KeyType]valueType

var map1 = map[string]int{"one": 1, "two": 2}

//map的删除与清空
// delete(map ,key)

func main() {
	delete(map1, "one")
	fmt.Print(map1)
}
