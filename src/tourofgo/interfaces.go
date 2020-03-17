package tourofgo

// I - expotable interface with method M() signature
type I interface {
	M()
}


func Interfaces() {
	// Interfaces
	/*
		1. Definition - Interface is a type, defines as a set of method signatures.
		2. A value of interface type can hold any method that implements its methods.
		3. Interfaces are implemented implicitly - so that implementation is in any package without any prearrangements.
		4. Interface value can be thought of as a tuple - (value, type)
		5. If the concrete value itself is null then the implementing method will be called with a nil reciever - in other languages it will throw null pointer exception.
		6. Empty interface type that specifies zero methods - interface{}.
		7. Empty interface can hold values of any type.
		8. Empty interfaces are used by code to handle values of unkown types.
	*/
	var i1 I = Vertex{5, 6} // Vertex type is implecitly implementing Interface type
	i1.M()
	var i2 I = F(3.4)
	i2.M()
	var f F
	var i3 I = f
	i3.M()

	// End of Interfaces

}