package algorithms

import (
	"fmt"
)

func FizzBuzz(n int) {

	for i:=1;i<n;i++ {
		printFizzBuzz(i)
		fmt.Print(", ")
	}
	printFizzBuzz(n)
	fmt.Println()
}

func printFizzBuzz(n int) {
	switch {
	case n%3==0 && n%5==0:
		fmt.Print("Fizz Buzz")
	case n%3==0:
		fmt.Print("Fizz")
	case n%5==0:
		fmt.Print("Buzz")
	default:
		fmt.Print(n)
	}
}