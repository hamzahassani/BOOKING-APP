package main

import "fmt"

func main() {

	var notes []int
	notes = append(notes, 1, 2, 3)
	for _, v := range notes {
		println(v)
	}

	m := map[string]int{
		"hello": 3,
		"test":  5,
	}

	v, ok := m["test"]
	fmt.Println(v, ok)
	for k, v := range m {
		println(k, v)
	}

}
