package main

import "fmt"

func main() {

	// slice

	var notes []int
	notes = append(notes, 1, 2, 3)
	for _, v := range notes {
		println(v)
	}

	// map
	var scores map[string]int
	fmt.Println(scores == nil)

	totalWins := map[string]int{}
	fmt.Println(totalWins == nil)

	m := map[string]int{
		"hello": 3,
		"test":  5,
	}

	v, ok := m["test"]
	fmt.Println(v, ok)
	for k, v := range m {
		fmt.Println(k, v)
		fmt.Println(k, v)
	}

	// struct
	type person struct {
		name string
		age  int
		pet  string
	}

	var fred person
	fmt.Println((fred))
	//A struct literal can be assigned to a variable as well:
	bob := person{"sayed", 42, "jaw"}
	fmt.Println((bob))
	beth := person{
		age:  30,
		name: "Beth",
	}
	fmt.Println((beth))

	sayed := person{}
	sayed.name = "sayed"
	fmt.Println((sayed))

}
