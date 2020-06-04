package main

import "fmt"

type Cat struct {
	Color string
	Name  string
}

func NewCatByName(name string) *Cat {
	return &Cat{
		Name: name,
	}
}

func NewCatByColor(Color string) *Cat {
	return &Cat{
		Color: Color,
	}
}

type BlackCat struct {
	Cat
}
type YeCat struct {
	Cat
}

//构造子类
func NewBlack(color string) *BlackCat {
	cat := &BlackCat{}
	cat.Color = color
	return cat
}

func NewYeCat(color string) *YeCat {
	cat := &YeCat{}
	cat.Color = color
	return cat
}

func main() {
	BlackCat := NewBlack("black")
	fmt.Println(BlackCat.Color)
}
