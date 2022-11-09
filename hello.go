package main

import "fmt"
import "integers"
const helloPrefix = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return helloPrefix + name
}

func main() {
	fmt.Println(Hello("Alex"))
	fmt.Println(len("Hello"))
	fmt.Println(Adder())
}
