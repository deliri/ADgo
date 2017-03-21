// BubbleSort has the worst runtime of O(n^2)
// Bubbles rise up and it takes a while, largest value moves far right
// Compare first element with the 2nd, if larger, swap
// Compare second element with third element, if larger, swap
// Keep going until the nth element is compared
// Repeat the whole process until the n-1th element.

package main

import (
	"fmt"
	"math/rand"
	"time"
)

// BubbleSort is the algo we use to see how long it takes
func BubbleSort(arr []int) []int {
	for i := len(arr) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if arr[j] > arr[j+1] {
				arr[j+1], arr[j] = arr[j], arr[j+1]
			}
		}
	}
	return arr
}

func main() {
	a := time.Now()
	var arr [100]int
	var slice = arr[:]

	// set up random num gen
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	// initialize the array with a random set of numbers from 0 to 10000000
	for index := range arr {
		arr[index] = random.Intn(1000000000)
	}
	// output before the sort
	fmt.Printf("Before: %v\n", arr)
	// sorting
	b := time.Now()
	BubbleSort(slice)
	// output after sort
	fmt.Printf("After: %v\n", arr)
	fmt.Printf("Time for sort to complete: %v\n", time.Since(b))
	fmt.Printf("Time overall: %v\n", time.Since(a))

}
