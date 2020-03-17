package tourofgo

import (
	"fmt"
)

func Collections() {
	//collections
	/* These are the only collections provided defaultly by GO
	Arrays
	Slices
	Map
	Structs
	*/

	// Array is an fixed size collection of similar datatype - statically sized.
	var arr [3]int // arrays can't be resized.
	arr[0] = 1

	arr2 := [3]int{1, 2, 3}
	fmt.Println("printing array", arr, arr2)
	// end of array

	//Slice data type - build on top of arrays - dynamically sized
	var sli = []int{1, 2, 3}
	fmt.Println("printing slice", sli)
	sli = append(sli, 4, 5, 6)
	fmt.Println("printing slice append method result", sli)

	// we can assign an array to slice (converting static to dyanamic - advantage of slice, can be usefull while working with api response)
	var slice = arr[:]
	fmt.Println("printing slice after assigning array to it", slice)
	slice = append(slice, 4, 5, 6, 7)
	arr[1] = 21
	slice[2] = 32
	slice[3] = 43
	fmt.Println("printing both array and slice", arr, slice)

	// slices os slices
	/* slices don't store any value it just refers to underlying array*/
	s2 := slice[1:]  // s2 starts with 1 and proceeds to end - includes position 1
	s3 := slice[:2]  // s3 starts with 0 and proceeds to 2 - not include position 2
	s4 := slice[1:2] // s4 starts with 1 and proceeds to 2 - includes position 1 but not 2
	s4[0] = 143      // update to slice will flow through other slices till parent(array).
	fmt.Println("printing slices of slices", arr, slice, s2, s3, s4)

	/* slice length and capacity
	1. length is the number of elements in an slice - len(slice).
	2. capacity is the number of elements in underlying array, counting from first element from slice - cap(slice).
	3. capacity will be effected when you drop any elements.
	4. s[2:] - this denotes droping the elements but not s[:4] =this.
	5. a := make([]int, 0 , 5) - make is the inbuilt function to create slice - len(a):0 , cap(a):5.
	6. range of a for loop can iterate over a slice.
	7. you can skip index or value in a range by assigning it to _ .
	8. if you only want index then you can skip the second parameter.
	*/

	//end of slice

	//Map
	/*
		1. map - is acombination of key, value pairs
		2. make function returns a map of given type, intialized and ready to use.
		3. zero value of a map is nil. nil map has 0 keys nor keys can be added.
		4. Intializing a map m := map[string]int {"x":1, "y":2}
		5. Insert an element - m[key] = ele
		6. retrieve an element - ele = m[key]
		7. delete an element - delete(m, key)
		8. test the key is present in a map or not - ele, ok := m[key] - if key is in m then ok=true or else ok=false.
		9. if element is not present in map then ele value is zero value of the map element type.
	*/
	// end of Map

	// Structs
	/*
		1. Definition - Struct type is a collection of fields.
		2. struct fields are accessed using a dot.
		3. struct fields can be accessed using a struct pointer - *p.x or p.x, where p is a struct.
	*/

	// Struct literals
	type vertex1 struct {
		x, y int
	}

	v1 := vertex1{3, 4}
	v2 := vertex1{x: 3}
	v3 := vertex1{}
	v4 := &vertex1{3, 4}
	fmt.Printf("printing struct literals v1:%v, v2:%v, v3:%v, v4:%v\n\n", v1, v2, v3, v4)

	// End of Structs

	// end of collecions

}