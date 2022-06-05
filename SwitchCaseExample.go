package main

import (
	"fmt"
	"reflect"
)

func main() {

	//declare a variable of type int
	var userInput int

	//Ask for the user input
	fmt.Print("Enter a number :")
	fmt.Scan(&userInput)

	//Print it to the terminal
	fmt.Println("You entered: ", userInput)
	fmt.Println("Type of userInput is: ", reflect.TypeOf(userInput))

	//determine if the entered input is even or odd using a switch statement
	switch userInput {
	case 0, 2, 4, 6, 8:
		fmt.Println("You entered Even Number")
	case 1, 3, 5, 7, 9:
		fmt.Println("You entered Odd Number")
	default:
		fmt.Println("Number is not in the range of 0-9")
	}

}
