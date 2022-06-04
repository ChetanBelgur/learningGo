package main

import "fmt"

func main() {
	var simpleInterest, principal, annualRateOfInterest float32
	var loanTenure int

	var amount float32

	fmt.Print("Enter Borrowed Loan Amount: ")
	fmt.Scan(&principal)

	fmt.Print("Enter Annual Rate Of Interest: ")
	fmt.Scan(&annualRateOfInterest)

	fmt.Print("Loan Tenure: ")
	fmt.Scan(&loanTenure)

	simpleInterest = (principal * annualRateOfInterest * float32(loanTenure)) / 100
	amount = simpleInterest + principal

	fmt.Println("Borrowed Amount: ", principal)
	fmt.Println("Simple Interest:", simpleInterest)
	fmt.Println("Total Amount: ", amount)

}
