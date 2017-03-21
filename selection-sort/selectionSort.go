// In computer science, selection sort is a sorting algorithm, specifically an in-place comparison sort. It has O(n2)
// time complexity, making it inefficient on large lists, and generally performs worse than the similar insertion sort.
// Selection sort is noted for its simplicity, and it has performance advantages over more complicated algorithms in
// certain situations, particularly where auxiliary memory is limited.
// The algorithm divides the input list into two parts: the sublist of items already sorted, which is built up from
// left to right at the front (left) of the list, and the sublist of items remaining to be sorted that occupy the rest
// of the list. Initially, the sorted sublist is empty and the unsorted sublist is the entire input list. The algorithm
// proceeds by finding the smallest (or largest, depending on sorting order) element in the unsorted sublist,
// exchanging (swapping) it with the leftmost unsorted element (putting it in sorted order), and moving the sublist
// boundaries one element to the right.

package main

import (
	"fmt"
	"math/rand"
	"time"
)

// SelectionSort implements the above algorithm
func SelectionSort(arr []int) []int {
	var min int
	var intHolder int
	for i := range arr {
		min = i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[i] {
				min = j
			}
		}
		// swap
		if min != i {
			intHolder = arr[i]
			arr[i] = arr[min]
			arr[min] = intHolder
		}
	}
	return arr
}

func main() {
	a := time.Now()
	var arr [1000]int
	var slice = arr[:]

	// set up random num gen
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	// initialize array with numbers from 0 to 10000000
	for index := range arr {
		arr[index] = random.Intn(10000000)
	}
	// output before sort
	fmt.Printf("Before: %v\n", arr)

	// sorting
	b := time.Now()
	result := SelectionSort(slice)

	// output after sort
	fmt.Printf("After: %v\n", result)
	fmt.Printf("Time for sort to complete: %v\n", time.Since(b))
	fmt.Printf("Time overall: %v\n", time.Since(a))

}
