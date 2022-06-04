package main

import "fmt"

func main() {
	for loopIndex := 0; loopIndex < 10; loopIndex++ {
		if loopIndex%2 == 0 {
			fmt.Print(loopIndex)
		} else {
			fmt.Print("\t")
			fmt.Println(loopIndex)
		}

	}
}
