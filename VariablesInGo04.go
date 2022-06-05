package main

import (
	"fmt"
	"reflect"
)

func main() {

	var student1 string = "Chetan" //Type is string
	var student2 = "Jane"          //type is inferred
	x := 2                         //type is inferred

	fmt.Println("Student1 is", student1)
	fmt.Println("Student2 is", student2)
	fmt.Println("x is", x)

	fmt.Println("Type of Student1 is", reflect.TypeOf(student1))
	fmt.Println("Type of Student2 is", reflect.TypeOf(student2))
	fmt.Println("Type of x is", reflect.TypeOf(x))
}
