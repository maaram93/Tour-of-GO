package tourofgo

import (
	"fmt"
)

func Methods() {
	// Methods
	/*
		1. we can define methods only on types.
		2. Definition - Method is a function with a special reciever argument.
		3. reciever argument should be defined in the same package.
		4. use pointers as a method arguments - advantages,
			1. method can modify the value that its reciever points to.
			2. efficient - avoids copying the value on each method call.
		5. Method should have either values or pointers as arguments but not both.
	*/

	v := Vertex{3, 4}
	mtd1 := v.firstMethod()
	fmt.Printf("printing values from firstMethod x: %v and y: %v\n", mtd1.x, mtd1.y)
	v.modifyRecieverMethod() // both lines 130 and 131 are good.
	//&v.modifyRecieverMethod() // ok -- only pointer reciever argument methods can modify reciever argument value but not the normal methods.
	mtd2 := v.firstMethod()
	fmt.Printf("printing values from modifyRecieverMethod x: %v and y: %v\n", mtd2.x, mtd2.y)
	// End of Methods

}