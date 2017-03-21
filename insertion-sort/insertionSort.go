// Insertion sort is a simple sorting algorithm that builds the final sorted array (or list) one item at a time. It is
// much less efficient on large lists than more advanced algorithms such as quicksort, heapsort, or merge sort.
// Insertion sort iterates, consuming one input element each repetition, and growing a sorted output list. Each
// iteration, insertion sort removes one element from the input data, finds the location it belongs within the sorted
// list, and inserts it there. It repeats until no input elements remain.
// Sorting is typically done in-place, by iterating up the array, growing the sorted list behind it. At each
// array-position, it checks the value there against the largest value in the sorted list (which happens to be next to
// it, in the previous array-position checked). If larger, it leaves the element in place and moves to the next. If
// smaller, it finds the correct position within the sorted list, shifts all the larger values up to make a space, and
// inserts into that correct position.
// source : https://en.wikipedia.org/wiki/Insertion_sort

package main

import (
	"fmt"
	"math/rand"
	"time"
)

// InsertionSort takes in a slice and gives out a slice sorted using the above algorithm
func InsertionSort(arr []int) []int {
	a := len(arr)
	for i := 1; i < a; i++ {
		for j := i - 1; j >= 0; j-- {
			if arr[j] > arr[j+1] {
				arr[j+1], arr[j] = arr[j], arr[j+1]
			} else {
				break
			}
		}
	}
	return arr
}

func main() {
	a := time.Now()
	var arr [100]int
	fmt.Println("arr: ", arr)
	var slice = arr[:]
	fmt.Println("slice: ", slice)

	// set up rand num
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	// initialize array with random from 0 to 10000000
	for index := range arr {
		arr[index] = random.Intn(10000000)
		fmt.Println("arr: ", arr)
	}

	// output before sort
	fmt.Printf("Before: %v\n", arr)
	// sorting
	b := time.Now()
	result := InsertionSort(slice)
	// output after sort
	fmt.Printf("After: %v\n", result)
	fmt.Printf("Time for sort to complete: %v\n", time.Since(b))
	fmt.Printf("Time overall: %v\n", time.Since(a))
}
