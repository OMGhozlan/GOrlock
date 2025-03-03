package main

import "fmt"

const englishHelloPrefix = "Hello, "

func HelloOld2() string {
	return "Hello, World!"
}
func HelloOld1(name string) string {
	return englishHelloPrefix + name
}
func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("world"))
}
