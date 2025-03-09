// main.go
package main

import (
	"fmt"

	"github.com/miraccan00/goodbye"
	"github.com/miraccan00/hello"
)

func main() {
	result1, result2 := hello.Hello("Miracle"), goodbye.Goodbye("Miracle")
	fmt.Println(result1)
	fmt.Println(result2)
}
