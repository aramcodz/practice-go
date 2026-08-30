package main

import "fmt"

func main() {
	fmt.Println(Hello("Benny"))
}

func Hello(name string) string {
	return "Hello, " + name
}
