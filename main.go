package main

import (
	"booking-app/closure"
	"errors"
	"fmt"
)

// named return values
func divAndRemainder(numerator int, denominator int) (result int, remainder int, err error) {
	if denominator == 0 {
		err = errors.New("cannot divide by zero")
		return result, remainder, err
	}
	result, remainder = numerator/denominator, numerator%denominator
	return result, remainder, err
}

//Functions Are Values

func main() {

	// slice

	/*var notes []int
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

	// Anonymous structs
	var persons struct {
		name string
		age  int
		pet  string
	}
	persons.name = "bob"
	persons.age = 50
	persons.pet = "dog"
	fmt.Println((persons))
	pet := struct {
		name string
		kind string
	}{
		name: "Fido",
		kind: "dog",
	}
	fmt.Println((pet))

	// comparing struct
	type firstPerson struct {
		name string
		age  int
	}
	f := firstPerson{
		name: "Bob",
		age:  50,
	}

	type secondPerson struct {
		name string
		age  int
	}

	g := secondPerson{
		name: "Bob",
		age:  50,
	}

	// compiles -- can use = and == between identical named and anonymous structs
	k := firstPerson(g)
	fmt.Println(k == f)
	fmt.Println(greetings.Hello("test"))

	// shadowed variables
	x := 10
	if x > 5 {
		fmt.Println(x)
		x := 10
		fmt.Println(x)
	}
	fmt.Println(x)
	var y int
	y = 13
	y, z := 15, 33
	fmt.Println(y, z)

	// if statement
	if n := rand.Intn(10); n == 0 {
		fmt.Println("That's too low")
	} else if n > 5 {
		fmt.Println("That's too big:", n)
	} else {
		fmt.Println("That's a good number:", n)
	}

	// named return values
	_, _, err := divAndRemainder(5, 2)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(err)

	// calculator
	calculator.Play()
	*/
	// Closure
	inc1 := closure.Createclosure(3)
	inc2 := closure.Createclosure(9)
	fmt.Println(inc1())
	fmt.Println(inc1())
	fmt.Println(inc2())
	fmt.Println(inc2())

}
