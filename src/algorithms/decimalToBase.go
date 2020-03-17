package algorithms

import (
	"fmt"
	"strconv"
)


func DecToBase(decimal, base int) {

	var output string
	var remainder int
	for decimal !=0 {
		remainder = decimal%base
		decimal = decimal/base
		output = strconv.Itoa(remainder) + output
	}
	fmt.Print(output)
}