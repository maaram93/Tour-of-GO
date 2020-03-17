package tourofgo

import (
	"fmt"
)

// pointers
	/*
		1. Definition - pointer holds the memory address of a value.
		2. type *T is a pointer to a T value, its zero value of a pointer is nil.
		3. & operator generates a pointer to its operand(quantity on which operation is to be done).
		4. * operator denotes underlying value of a pointer.
	*/

	func Pointers() {

	lastName := "MRM"
	ptr := &lastName // pointer for lastName
	fmt.Println("variable and coresponding pointer (memory address)", lastName, &lastName, ptr, *ptr)
	lastName = "Reddy"
	fmt.Println("updated value but pointer (memory address) remains same", lastName, &lastName)

	// declare and Intializa a pointer
	var ptr1 *string = new(string)
	*ptr1 = "Raj" // dereferencing the pointer and asssign a value
	fmt.Println("printing memory address of ptr1", ptr1, *ptr1)

	}

// end of pointers
