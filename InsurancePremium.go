package main

import "fmt"

const (
	DAYS_IN_A_YEAR  = 365
	DAILY_AMOUNT    = 15
	NUMBER_OF_YEARS = 25
)

func main() {
	fmt.Println("Your Premium Amount for", NUMBER_OF_YEARS, "years is", DAYS_IN_A_YEAR*DAILY_AMOUNT*NUMBER_OF_YEARS)
}
