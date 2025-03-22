package main

//run go mod tidy
//When you ran go mod tidy, it located and downloaded the rsc.io/quote/v4 module that contains the package you imported.
//import the fmt package
//import the quote package
//create a main function
//call the fmt.Println function and pass the quote.Go() function as an argument
//run the program
//go run hello.go
//output
//Don't communicate by sharing memory, share memory by communicating.
import (
	"fmt"

	"rsc.io/quote/v4"
)

func main() {
	fmt.Println(quote.Go())
}
