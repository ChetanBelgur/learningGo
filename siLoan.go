package main

import "fmt"

func main() {

	//Declare Variables to compute
	var simpleInterest, principal, annualRateOfInterest float32
	var loanTenure int

	var amount float32

	//Ask User Inputs
	fmt.Print("Enter Borrowed Loan Amount: ")
	fmt.Scan(&principal)

	fmt.Print("Enter Annual Rate Of Interest: ")
	fmt.Scan(&annualRateOfInterest)

	fmt.Print("Loan Tenure: ")
	fmt.Scan(&loanTenure)

	//Compute simple Interest and final amount to pay
	simpleInterest = (principal * annualRateOfInterest * float32(loanTenure)) / 100
	amount = simpleInterest + principal

	//Print computed values
	fmt.Println("Borrowed Amount: ", principal)
	fmt.Println("Simple Interest:", simpleInterest)
	fmt.Println("Total Amount: ", amount)

}
