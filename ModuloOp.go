package main

import (
	"fmt"
	"reflect"
)

func main() {

	//Declare a variable on type int
	var numInput int

	//Ask to enter an input from the user
	fmt.Print("Enter a number: ")
	fmt.Scan(&numInput)

	//Print the entered number and its type to the terminal
	fmt.Println("You entered: ", numInput)
	fmt.Println("Type of numInput is:", reflect.TypeOf(numInput))

	//Check if its an odd or even number using the modulo operator
	if numInput%2 == 0 {
		fmt.Println("You entered an even Number")
	} else {
		fmt.Println("You entered an Odd Number")
	}
}
