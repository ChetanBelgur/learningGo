package main

//Packages
import (
	"fmt"
	"reflect"
	"strconv"
)

//Package level variables
var (
	name    = "chetan belgur"
	age     = "45"
	place   = "Bangalore"
	country = "India"
	height  = 6
)

func main() {

	//Print the to the terminal
	fmt.Println("Name: ", name)
	fmt.Println("Age: ", age)
	fmt.Println("Place:", place)
	fmt.Println("Country: ", country)

	//Get their types
	fmt.Println("Type of Name: ", reflect.TypeOf(name))
	fmt.Println("Type of Age: ", reflect.TypeOf(age))

	// Sum up name and age
	//total := height + age
	iAge, err := strconv.Atoi(age)
	if err == nil {
		fmt.Println("Type of Age after conversion is:", reflect.TypeOf(iAge))
		total := iAge + height
		fmt.Println("Total is:", total)
	}

}
