package main

import (
	"fmt"
	"reflect"
)

//Package level variables
var (
	name    = "Chetan Belgur"
	age     = 45
	place   = "Bangalore"
	Country = "India"
)

func main() {

	//Print them to the terminal
	fmt.Println("Name is:", name)
	fmt.Println("Age is:", age)
	fmt.Println("Place is:", place)
	fmt.Println("Country is:", Country)

	//Print the type of the variables
	fmt.Println("Type of Name is:", reflect.TypeOf(name))
	fmt.Println("Type of Age is:", reflect.TypeOf(age))

}
