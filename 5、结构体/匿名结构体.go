package main

import "fmt"

type name struct {
}

func printMsg(msg *struct {
	id   int
	data string
}) {
	fmt.Println(msg.data)
}

func main() {
	msg := struct {
		id   int
		data string
	}{
		1024,
		"hello",
	}

	printMsg(&msg)
}
