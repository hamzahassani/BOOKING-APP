package calculator

import (
	"errors"
	"fmt"
	"strconv"
)

func add(a int, b int) int {
	return a + b
}

func sub(a int, b int) int {
	return a - b
}
func mul(a int, b int) int {
	return a * b
}
func div(a int, b int) int {
	return a / b
}

func createMap() map[string]func(int, int) int {
	Operations := map[string]func(int, int) int{
		"+": add,
		"-": sub,
		"*": mul,
		"/": div,
	}
	return Operations

}

func Play() {
	expressions := [][]string{[]string{"2", "+", "4"}, []string{"d", "-", "4"}, []string{"2", "*", "4"}, []string{"22", "/", "0"}}
	for _, expression := range expressions {
		if len(expression) != 3 {
			fmt.Println("invalid expression:", expression)
			continue
		}
		n, err := strconv.Atoi(expression[0])
		if err != nil {
			fmt.Println(err)
			continue
		}
		op := expression[1]
		opMap := createMap()
		opFunc, ok := opMap[op]
		if !ok {
			fmt.Println("unsupported operator:", op)
			continue
		}
		d, err := strconv.Atoi(expression[2])
		if err != nil {
			fmt.Println(err)
			continue
		}
		if expression[1] == "/" {
			if d == 0 {
				err := errors.New("invalid opeartion! divide by zero ")
				fmt.Println(err)
				continue
			}
		}

		fmt.Println(opFunc(n, d))

	}

}
