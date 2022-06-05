package main

import (
	"fmt"
	"reflect"
	"strings"
)

func main() {

	var age int
	var firstName string
	var lastName string

	fmt.Print("Enter your Age: ")
	fmt.Scan(&age)
	fmt.Println("You Entered: ", age)
	fmt.Println("Type of num1 is:", reflect.TypeOf(age))

	fmt.Print("Enter your First Name: ")
	fmt.Scan(&firstName)

	fmt.Print("Enter Your Last Name: ")
	fmt.Scan(&lastName)

	fmt.Println("Type of First Name is", reflect.TypeOf(firstName), "and Last Name is:", reflect.TypeOf(lastName))

	firstName = strings.ToUpper(firstName)
	lastName = strings.ToUpper(lastName)
	fmt.Println("Hi,", firstName+" "+lastName, "You are already", age, "Years old...")

}
