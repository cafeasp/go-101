package main

import "fmt"

func main() {
	// variable declaration
	//var firstName = "bob"
	//var lastName = "smith"
	//fmt.Println("Hello, my name is", firstName, lastName)

	// shorthand variable declaration
	//middleName := "joe"
	//fmt.Println("Hello, my name is", firstName, middleName, lastName)

	//ints
	var age int = 40
	fmt.Println("I am", age, "years old")
	oldAge := 100
	fmt.Println("I will be", oldAge, "years old")

	//bits
	var b int8 = 13
	fmt.Println("b is", b)
	var c int16 = 2
	fmt.Println("c is", c)

	//floats
	var f float32 = 3.14
	fmt.Println("f is", f)
	var g float64 = 6.283185
	fmt.Println("g is", g)
}
