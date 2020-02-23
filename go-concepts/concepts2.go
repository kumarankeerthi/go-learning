package main

import (
	"fmt"
)

// *** STRUCT ****

type person struct {
	name string
	age  int
}

func main() {
	fmt.Println("Inside main func")
	fmt.Println("Calling sum function ", sum(4, 5))
	fmt.Println("Calling multi return func - have to accept return value as vars")
	i, s, b := multi_return_func()
	fmt.Println("return values are", i, s, b)

	fmt.Println("invoking Variadic func")
	variadicfunc(1, 2, 3, 4, 5)
	variadicfunc(1, 2)
	nums := []int{1, 2, 4, 5, 6}
	variadicfunc(nums...)

	// Ananymous function

	func(msg string) {
		fmt.Println("Hellos", msg)
	}("Golang")

	printer := printer_analymous()

	fmt.Println("invovking printer func")
	printer("Golang is simple yet complex")

	fmt.Println("Clouser example")

	counter := next_counter()
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println("reseting counter")

	counter = next_counter()
	fmt.Println(counter())

	fmt.Println("Structs and Pointer examplse")

	p1 := person{name: "Jon", age: 30}
	fmt.Println(p1, p1.name, p1.age, &p1)
	p2 := newPerson("Akki")
	fmt.Println(p2, p2.name, p2.age, &p2, *p2)

	fmt.Println(p1)
	fmt.Println(p2)

	p1.printUsingPointer()
	p2.printUsingPointer()
	p1.printUsingPointer()
	p2.printUsingPointer()
	//p1.printUsingValue()

}

// ******* FUNCTIONS *******

func sum(a int, b int) int {
	// func sum(a,b int) int
	return a + b
}

func multi_return_func() (int, string, bool) {
	return 9, "golang", true
}

func variadicfunc(nums ...int) {
	for _, num := range nums {
		fmt.Println(num)
	}
}

func printer_analymous() func(string) {
	return func(msg string) {
		fmt.Println("Printing input msg:", msg)
	}
}

func next_counter() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 32
	return &p
}

// ******* METHODS ******

func (p *person) printUsingPointer() {
	p.age += 1
	fmt.Println(p.name, p.age)
}

/*func (p person) printUsingValue() {
	fmt.Println(p.name, p.age)
}

func (p *person) String() string {
	return fmt.Sprintln("Pretty printing person :", p.name, p.age)
}*/
