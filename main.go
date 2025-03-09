package main

import (
	"fmt"

	"github.com/miraccan00/learn-golang-modules/goodbye"
	"github.com/miraccan00/learn-golang-modules/greeting"
	"github.com/miraccan00/learn-golang-modules/hello"
)

func main() {
	fmt.Println(hello.Hello("Miracle"))
	fmt.Println(goodbye.Goodbye("Miracle"))
	fmt.Println(greeting.Greet("Miracle", true))
	fmt.Println(greeting.Greet("Miracle", false))
}
