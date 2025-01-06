package main

import "fmt"

func main() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	i = 0.43

	f, ok := i.(float64)
	fmt.Println(f, ok)

	switch i.(type) {
	case float64:
		fmt.Println("float", i)
	case string:
		fmt.Println("string", i)
	}
}
