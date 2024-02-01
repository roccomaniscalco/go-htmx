package main

import (
	"fmt"
	// import from moduleName/packageName
	"mody/stack"
)

// var can be accessed anywhere in the package
var globalString = "Hello, World!"

func main() {
	helloWorld()
	printPyramid(11)

	// := operator assigns value to a locally scoped variable
	// only accessible within the scope it is defined
	s := stack.Stack[int]{}

	s.Push(3)
	var element = s.Pop()
	fmt.Println(element)
}

func helloWorld() {
	fmt.Println(globalString)
}

func printPyramid(height int) {
	for i := 0; i < height; i++ {
		for j := 0; j < height*2; j++ {
			if j < height-i || j > height+i {
				fmt.Print(" ")
			} else {
				fmt.Print("*")
			}
		}
		fmt.Println()
	}
}
