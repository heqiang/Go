package main

import "fmt"

type Humaner interface {
	SayHi()
}

type Student struct {
	name string
	id   int
}

func (temp *Student) SayHi() {
	fmt.Printf("Student%s %d say hi\n", temp.name, temp.id)
}

type str_ string

func (str *str_) SayHi() {
	fmt.Printf("MyStr %s sayhi\n", *str)
}

func main() {

	stu := &Student{"mike", 1}
	stu.SayHi()
	fmt.Println(stu.id)

	var sss str_ = "hqhqhqh"
	t := &sss
	t.SayHi()

}
