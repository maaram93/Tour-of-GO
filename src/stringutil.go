package main

import (
	"fmt"
	"stringutil"
)

func main() {
	fmt.Println(stringutil.Reverse("hello go! Version 1.0"))
	fmt.Println(stringutil.ReverseVersion2("hello go! Version 2.0"))
	fmt.Println(stringutil.ReverseVersion3("hello go! Version 3.0"))
	stringutil.WordCount("I ate a donut. Then I ate another donut.")
}