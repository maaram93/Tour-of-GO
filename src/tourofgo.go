package main

import (
	"fmt"
	"tourOfGO"
	"runtime"
)

func main() {
	
	fmt.Println("hello go")
	fmt.Println("Hello from", runtime.GOOS) 

	// variables
	fmt.Println("name and course are set to", tourOfGO.Name)
	fmt.Println("module is set to", tourOfGO.Module)
	tourOfGO.Variables();

	// pointers
	tourOfGO.Pointers();

	// collections
	tourOfGO.Collections();

	// methods
	tourOfGO.Methods();

}
