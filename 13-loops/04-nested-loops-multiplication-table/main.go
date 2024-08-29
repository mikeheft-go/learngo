// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type operation func(int, int) int

func add(a, b int) int {
	return a + b
}

func subtract(a, b int) int {
	return a - b
}

func multiply(a, b int) int {
	return a * b
}

func divide(a, b int) int {
	if b == 0 {
		panic("division by zero")
	}
	return a / b
}

func main() {
	// print the header
	args := os.Args[1:]
	if len(args) < 2 || len(args) > 2 {
		panic("[ERROR] Usage: go run main.go <operator(x)> <number(5)>")
	}

	operator := args[0]
	matched, _ := regexp.MatchString(`^[x\-+\/]$`, operator)
	fmt.Println(args)
	if len(operator) > 1 || !matched {
		fmt.Printf("[ERROR] Valid operators are %[1]q, %[2]q, %[3]q, %[4]q", "x", "-", "+", "/")
	}

	max, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println(err)
		panic(nil)
	}

	fmt.Printf("%5s", operator)
	for i := 0; i <= max; i++ {
		fmt.Printf("%5d", i)
	}
	fmt.Println()

	ops := map[string]operation{
		"+": add,
		"-": subtract,
		"*": multiply,
		"/": divide,
	}

	for i := 0; i <= max; i++ {
		// print the vertical header
		fmt.Printf("%5d", i)

		// print the cells
		for j := 0; j <= max; j++ {
			fmt.Printf("%5d", ops[operator](i, j))
		}
		fmt.Println()
	}
}
