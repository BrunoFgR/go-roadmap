package main

import "fmt"

type Dog struct {
	name  string
	age   int
	sound string
}

func (d *Dog) Sound() string {
	return d.sound
}

type Cat struct {
	name  string
	age   int
	sound string
}

func (c *Cat) Sound() string {
	return c.sound
}

type Animals interface {
	Sound() string
}

func main() {
	var animal Animals

	dog := &Dog{name: "Duck", age: 12, sound: "Au Au"}

	animal = dog
	fmt.Println("animal sound:", animal.Sound())

	cat := &Cat{name: "Pantera", age: 2, sound: "Miau Miau"}
	animal = cat

	fmt.Println("animal sound:", animal.Sound())
}
