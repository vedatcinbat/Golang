package main

import "fmt"

// functions
func add(a int, b int) int {
	return a + b
}

func main() {
	// constants
	/* const pi = 3.14159
	const appName = "GoLearning" */

	// using var
	/* var num1 int = 10
	var b = 20

	c := 30

	fmt.Println(num1, b, c)
	fmt.Println("App Name:", appName)
	fmt.Println("Value of Pi: ", pi)

	result := add(10, 20)
	fmt.Println("Result:", result) */

	// Conditionals
	/* x := 10
	if x > 5 {
		fmt.Println("x is greater than 5")
	} else {
		fmt.Println("x is less than 5")
	} */

	/* day := "Monday"

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
	} */

	// Arrays
	/* var arr [3]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3

	fmt.Println(arr)

	arr2 := [3]string{"Go", "Rust", "Python"}
	arr3 := [...]int{1, 2, 3, 4, 5}

	fmt.Println(arr2)
	fmt.Println(arr3)

	// Slices
	slice := []int{10, 20, 30, 40, 50}

	fmt.Println(slice)

	// Create a slice from an array
	arr4 := [5]int{1, 2, 3, 4, 5}
	slice2 := arr4[1:4] // indexes 1 to 3 (not 4)

	fmt.Println(slice2) */

	// Append items
	/* numbers := []int{1, 2, 3}
	numbers = append(numbers, 4, 5)

	fmt.Println("Appended Array:", numbers)
	fmt.Println("Length:", len(numbers))
	fmt.Println("Capacity:", cap(numbers)) */

	// Maps
	/* m := make(map[string]int)
	m["apple"] = 10
	m["banana"] = 20

	fmt.Println(m)*/

	// Inline initialization
	m := map[string]string{
		"name": "Vedat Cinbat",
		"city": "Istanbul",
	}

	fmt.Println(m)
	fmt.Println("Name", m["name"])

	v, ok := m["apple"]
	if ok {
		fmt.Println("Apple exists with value", v)
	} else {
		fmt.Println("Apple not found")
	}

}
