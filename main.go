package main

import "fmt"

func main() {

	//Print
	// fmt.Print("Hello, ")
	// fmt.Print("World")

	// fmt.Println("One line")
	// fmt.Println("Next line")

	// printing variables
	name := "Eddy"
	age := "27"
	numt := 232
	//fmt.Println("My name is", name, ",I am", age, "years old")

	// string formatting // %_ = format specifier
	fmt.Printf("My age is %v and my name is %v \n", name, age)
	fmt.Printf("My age is %q and my name is %q \n", name, age)
	fmt.Printf("numt is of type %T \n", numt)
	fmt.Printf("You scored %0.2f points! \n", 223.43)
}
