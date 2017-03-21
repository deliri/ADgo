// Array and slices in go
package main

import "fmt"

func main() {
	fmt.Printf("Arrays and Slices in go.\n")
	var arrayA [10]int = [10]int{1, 2, 3, 4, 5} // literal array
	fmt.Println("arrayA is: ", arrayA)
	var sliceA []int // slice
	fmt.Println("sliceA is: ", sliceA)
	fmt.Printf("The type of the Array is %T while type of slice is %T\n", arrayA, sliceA)
	fmt.Printf("Appear to be the same, arrays are just a list of elements, while slices are a view into the array.\n")

	// Make a slice out of an Array
	sliceA = arrayA[0:9] // sliceA has a full view of arrayA
	fmt.Printf("This is sliceA %v\n", sliceA)

	var sliceB []int
	fmt.Printf("This is sliceB %v\n", sliceB)
	sliceB = arrayA[1:9] // sliceB has a view of the index of the array minus the 0 index
	fmt.Printf("This is sliceB taken from arrayA[1:9] %v\n", sliceB)

	fmt.Printf("Lets see the length len() of sliceA: %v and sliceB: %v\n", len(sliceA), len(sliceB))
	fmt.Printf("We can also see the capacity of the slices.\n")
	fmt.Printf("Capacity of sliceA: %v and sliceB: %v\n", cap(sliceA), cap(sliceB))

	// change the array data
	fmt.Printf("Lets change the data inside the array.\n")
	arrayA[2] = 1000 // this will show up in the sliceA and sliceB
	// sliceB started at index 1 of arrayA and sliceB[1] == arrayA[2]
	fmt.Printf("Value of sliceA[2]: %v and sliceB[1]: %v\n", sliceA[2], sliceB[1])
	// Lets change some slice data and see how that shows up
	fmt.Printf("Changing a slice will access the array and change the corresponding data\n")
	sliceA[2] = 3000
	fmt.Printf("What are the new values of arrayA[2]: %v and sliceB[1]: %v\n", arrayA[2], sliceB[1])

}
