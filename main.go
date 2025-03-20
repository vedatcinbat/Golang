package main

import "fmt"

// functions
func add(a int, b int) int {
	return a + b
}

func main() {
	// constants
	const pi = 3.14159
	const appName = "GoLearning"

	// using var
	var num1 int = 10
	var b = 20

	c := 30

	fmt.Println(num1, b, c)
	fmt.Println("App Name:", appName)
	fmt.Println("Value of Pi: ", pi)

	result := add(10, 20)
	fmt.Println("Result:", result)

	// Conditionals
	x := 10
	if x > 5 {
		fmt.Println("x is greater than 5")
	} else {
		fmt.Println("x is less than 5")
	}

	day := "Monday"

	switch day {
	case "Monday":
		fmt.Println("Today is Monday")
	case "Tuesday":
		fmt.Println("Today is Tuesday")
	default:
		fmt.Println("Just another day")
	}

	// Loops
	// For loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// While loop
	s := 0
	for s < 5 {
		fmt.Println(s)
		s++
	}
}

/*
* Common Data Types
	var i int = 42
	var f float64 = 3.14
	var s string = "Hello, Go"
	var b bool = true

* Default Zero Values
	var i int // 0
	var f float64 // 0
	var s string // ""
	var b bool // false

* Conditionals

if x > 5 {
	fmt.Println("x is greater than 5")
}



*/
