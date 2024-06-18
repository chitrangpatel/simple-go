package main

import (
	"fmt"

	"github.com/chitrangpatel/simple-go/pkg/math"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Printf("1 + 3 = %d\n", math.Add(1, 3))
	fmt.Printf("3 - 1 = %d\n", math.Subtract(3, 1))
}
