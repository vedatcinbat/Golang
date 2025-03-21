/* package main

import "fmt" */

// functions
/* func add(a int, b int) int {
	return a + b
}

type Person struct {
	Name string
	Age  int
	City string
}

func (p Person) Greet() {
	fmt.Println("Hello, my name is ", p.Name, "and I am ", p.Age, "years old.")
}

func (p *Person) Birthday() {
	p.Age++
}

func main() { */
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
/* m := map[string]string{
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
} */

// Structs (Objects)
/* p1 := Person{
	Name: "Vedat", Age: 30, City: "Istanbul",
}

p2 := Person{"John", 25, "New York"}

var p3 Person
p3.Name = "Alice"
p3.Age = 35
p3.City = "London"

fmt.Println(p1)
fmt.Println(p2)
fmt.Println(p3)

p := &Person{Name: "Diana", Age: 28}
fmt.Println(p.Name) */

//p := Person{Name: "Emily", Age: 24}
//p.Greet()
//
//p.Birthday()
//p.Greet()
//}

// Goroutines
/* package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Hello from Goroutine !")
}

func main() { */
/* go sayHello() // Runs concurrently with the main function

fmt.Println("Main function running")
time.Sleep(100 * time.Millisecond) */

/* ch := make(chan string)

go func() {
	ch <- "Hello from Goroutine !"
}()

msg := <-ch
fmt.Println(msg) */

/* ch := make(chan int, 2)
ch <- 10
ch <- 20

fmt.Println(<-ch)
fmt.Println(<-ch) */

/* ch := make(chan int)

go func() {
	for i := 1; i <= 5; i++ {
		ch <- i
	}
	close(ch)
}()

for val := range ch {
	fmt.Println(val)
} */

/* ch1 := make(chan string)
ch2 := make(chan string)

go func() {
	time.Sleep(1 * time.Second)
	ch1 <- "message from ch1"
}()

go func() {
	time.Sleep(3 * time.Second)
	ch2 <- "message from ch2"
}()

for i := 0; i < 2; i++ {
	select {
	case msg1 := <-ch1:
		fmt.Println(msg1)
	case msg2 := <-ch2:
		fmt.Println(msg2)
	}
} */

//}

// Making HTTP Requests and Handling Responses
/* package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

func fetchUser(id int, wg *sync.WaitGroup, ch chan string) {
	defer wg.Done()

	url := fmt.Sprintf("https://jsonplaceholder.typicode.com/users/%d", id)
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("Error fetching user %d: %s", id, err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("Error reading response body: %s", err)
		return
	}

	// Send response body as string through the channel
	ch <- fmt.Sprintf("User %d: %s", id, body)
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan string)

	userIDs := []int{1, 2, 3, 4, 5}

	for _, id := range userIDs {
		wg.Add(1)
		go fetchUser(id, &wg, ch)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for result := range ch {
		fmt.Println(result)
	}
} */

// Go Turorial
package main

import (
	"fmt"
)

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
	/* fmt.Println("My fav nmber is:", rand.Intn(10))

	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}

	fmt.Println("Sum:", sum) */

	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)

	x, y := <-c, <-c

	fmt.Println(x, y, x+y)

	c2 := make(chan int, 10)
	go fibonacci(cap(c2), c2)
	for i := range c2 {
		fmt.Println(i)
	}

}
