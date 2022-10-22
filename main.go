package main

import (
	"fmt"
)

func variables() {
	var i int = 42
	j := 23
	var k int
	k = 68

	fmt.Println("Hello World!")
	fmt.Printf("Number i is %v.\n", i)
	fmt.Printf("Number j is %v.\n", j)
	fmt.Printf("Number k is %v.\n", k)
}

func arrays() {
	grades := [3]int{93, 79, 84}
	grades_b := [...]int{93, 43, 60}

	grades_c := &grades

	fmt.Printf("Grade: %v\n", grades)
	fmt.Printf("Grade: %v\n", grades_b)

	grades_c[1] = 99
	fmt.Printf("Grade: %v\n", grades_c)
	fmt.Printf("Original grade: %v\n", grades)
}

func main() {
	variables()
	arrays()
}
