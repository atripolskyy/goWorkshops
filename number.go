package main

import (
	"fmt"
	"math/rand"
)

func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println("My favorite number is", rand.Intn(add(12, 11)))
}
