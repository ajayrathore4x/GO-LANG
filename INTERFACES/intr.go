package main

import "fmt"

// 1. Interface
type Speaker interface {
	Speak()
}

// 2. Dog struct
type Dog struct {
	name string
}

// Dog has Speak(), so Dog satisfies Speaker
func (d Dog) Speak() {
	fmt.Println(d.name, "says Woof")
}

// 3. Human struct
type Human struct {
	name string
}

// Human also has Speak(), so Human satisfies Speaker
func (h Human) Speak() {
	fmt.Println(h.name, "says Hello")
}

// 4. This function accepts anything that satisfies Speaker
func makeSound(s Speaker) {
	s.Speak()
}

func main() {

	// Create a Dog
	dog := Dog{name: "Bruno"}

	// Create a Human
	human := Human{name: "Ajay"}

	// Both can be passed to makeSound()
	makeSound(dog)
	makeSound(human)
}