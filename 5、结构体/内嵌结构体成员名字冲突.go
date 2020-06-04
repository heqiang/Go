package main

type A struct {
	a int
}

type B struct {
	a int
}

type C struct {
	A
	B
}

func main() {
	SS := C{}
	SS.A.a = 1

}
