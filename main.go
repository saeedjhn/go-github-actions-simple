package main

import (
	"fmt"
)

func main() {
	msg := sayHello("Alice")
	fmt.Println(msg)
}

func sayHello(name string) string {
	return fmt.Sprintf("Hi %s", name)
}
