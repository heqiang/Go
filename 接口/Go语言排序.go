package main

import (
	"fmt"
	"sort"
)

type MyStringList []string

func (m MyStringList) Len() int {
	return len(m)
}

func (m MyStringList) Less(i, j int) bool {
	return m[i] < m[j]
}

func (m MyStringList) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func main() {
	names := MyStringList{
		"5 Tr",
		"2 ts",
		"3 hq",
		"7,dd",
	}
	sort.Sort(names)
	for _, v := range names {
		fmt.Println(v)
	}
}
