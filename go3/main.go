package main

import (
	"fmt"
)

type Person interface {
	greet() string
}

type Human struct {
	Name string
}

func (h *Human) greet() string {
	return "Hi, I am " + h.Name
}

type Dog struct {
}

func (d *Dog) bark() string {
	return "Woof"
}

func say(h Person) {
	fmt.Println(h.greet())
}

func main() {
	var a = Human{"John"}
	var d = Dog{}

	fmt.Println(a.greet()) // Hi, I am John
	fmt.Println(d.bark())  // Woof

	// below function will only work
	// if `a` is also a person.
	// Here we can see polymorphism in action.
	say(&a) // Hi, I am John
	// say(&d) // won't compile
}
