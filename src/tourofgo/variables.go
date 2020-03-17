package tourofgo

import (
	"fmt"
)

var (
	Name, course string // declaring a variable
	Module       float64
	version      = "1.0" // declaring and intializing a variable
)

const (
	/* iota is only valid for package level variables - iota are atomatically increment values on each execution staring from 0 - iota are compile time*/
	first = iota
	second
)

const (
	third = iota
	fourth
)

// Vertex - exportable variable of type struct
type Vertex struct {
	x, y float64
}

// F - exportable variable of type float64
type F float64

func (v Vertex) firstMethod() Vertex {
	return v
}

func (v *Vertex) modifyRecieverMethod() {
	v.x = v.x * v.x
	v.y = v.y * v.y
}

// M - of type Vertex implictly implementing interface I
func (v Vertex) M() {
	fmt.Printf("Printing vertex type variable using interface impl M():  x: %v, y: %v and type %T\n\n", v.x, v.y, v)
}

// M - of type Vertex implictly implementing interface I
func (f F) M() {
	fmt.Printf("Printing vertex type variable using interface impl M():  x: %v and type %T\n\n", f, f)
}

// Variables - sample function
func Variables() {

	// self excuting function - f() {}() - the ending paranthesis make a function self excutable.

	// Type Conversion
	// the expression T(v) converts the value v to type T
	var intgr = 100
	var flot = float64(intgr)
	var u = uint(flot)
	fmt.Println("Type conversion", intgr, flot, u)
	// End of type convertion

	// implicit variable are allowed in methods only
	firstName := "raj" //Implicit Initialization of a variable
	fmt.Println("Implicit variable", firstName)
	c := complex(3, 4) // declaring a complex number
	fmt.Println("complex data type variable c", c)
	// spliting a complex variable
	r, i := real(c), imag(c)
	fmt.Printf("spliting complex variable real value %v and Imaginary value %v\n", r, i)

	// constants
	/* IMPORTANT - value of constant should be able to determined at compile time
	can't assign function return to a constant as functions are evaluated runtime.
	*/
	const pi = 3.145 // Implicit type constant compiler will interpret the type every it runs
	//pi = 5.143 // cant reassign a value to constant and should be intialized at declaration time only.
	fmt.Println("constants", pi)
	// end of constants

	// constant expressions
	fmt.Println("printing package level iota constants", first, second, third, fourth)
	// end of constant expressions

}
