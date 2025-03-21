package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}

//before running the code, you need to install Go on your machine. You can download it from the official website of Go.
//run go mod init hello
//This will create a new module named hello in the current directory. A module is a collection of Go packages stored in a file tree with a go.mod file at its root.
//The go.mod file defines the module's path, which is the import path used by other packages to import the module. In this case, the module path is hello.
//Next, you need to create a new Go file named hello.go in the same directory and add the following code to it:
//The code defines a main function that prints "Hello, World!" to the standard output using the fmt.Println function.
//The fmt package is part of the Go standard library and provides functions for formatting and printing text.
//To run the code, you need to open the terminal and navigate to the directory where the code is saved. Then, you can run the code using the following command:
//go run hello.go
//This will compile and run the code, and you should see the output "Hello, World!" printed on the terminal.
//keep the mod file with project and add the go file in the same directory
