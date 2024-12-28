package main

import "fmt"

func main() {
	f()
	fmt.Println("Returned normally from f")
	slice := []int{1, 2, 3, 4, 5}
	pop(slice)
	fmt.Println(slice)
}

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	fmt.Println("Calling in g")
	g(0)
	fmt.Println("Returned normally from g")
}

func pop(slice []int) {
	slice2 := append(slice[:2], slice[3:]...)
	fmt.Println(slice2)
}

func g(value int) {
	if value > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%d", value))
	}
	defer fmt.Println("Defer", value)
	fmt.Println("Printing in g", value)
	g(value + 1)
}
