package main

import (
	"algorithms"
	"fmt"
	"bufio"
	"os"
	"strconv"
)

func main() {
	
	// Fizz Buzz problem
	fmt.Print("Enter Fizz_Buzz input number:")
	var input int
	fmt.Scanln(&input)
	algorithms.FizzBuzz(input)

	// Decimal to Base problem
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter Decimal number:")
	scanner.Scan()
	input1 := scanner.Text()
	decimal,_ := strconv.Atoi(input1)
	fmt.Print("Enter Base:")
	scanner.Scan()
	input2 := scanner.Text()
	base,_ := strconv.Atoi(input2)
	algorithms.DecToBase(decimal,base);
	
}