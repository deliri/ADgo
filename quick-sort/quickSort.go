// Quicksort (sometimes called partition-exchange sort) is an efficient sorting algorithm, serving as a systematic
// method for placing the elements of an array in order. Developed by Tony Hoare in 1959,[1] with his work published in
// 1961, it is still a commonly used algorithm for sorting. When implemented well, it can be about two or three times
// faster than its main competitors, merge sort and heapsort.[3]
// Quicksort is a comparison sort, meaning that it can sort items of any type for which a "less-than" relation (
// formally, a total order) is defined. In efficient implementations it is not a stable sort, meaning that the relative
// order of equal sort items is not preserved. Quicksort can operate in-place on an array, requiring small additional
// amounts of memory to perform the sorting.
// Mathematical analysis of quicksort shows that, on average, the algorithm takes O(n log n) comparisons to sort n
// items. In the worst case, it makes O(n2) comparisons, though this behavior is rare.

// Quicksort is a divide and conquer algorithm. Quicksort first divides a large array into two smaller sub-arrays: the
// low elements and the high elements. Quicksort can then recursively sort the sub-arrays.
// The steps are:
// Pick an element, called a pivot, from the array.
// Partitioning: reorder the array so that all elements with values less than the pivot come before the pivot, while
// all elements with values greater than the pivot come after it (equal values can go either way). After this
// partitioning, the pivot is in its final position. This is called the partition operation.
// Recursively apply the above steps to the sub-array of elements with smaller values and separately to the sub-array
// of elements with greater values.
// The base case of the recursion is arrays of size zero or one, which never need to be sorted.
// The pivot selection and partitioning steps can be done in several different ways; the choice of specific
// implementation schemes greatly affects the algorithm's performance.
// source : https://en.wikipedia.org/wiki/Quicksort

package main

import (
	"fmt"
	"math/rand"
	"time"
)

// swap initiates the swaps when comparing two elements in an array
func swap(arr []int, i int, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

// quickSort takes in a slice and using a pivot point recursively sorts the data and returns a slice
func quickSort(arr []int, left int, right int) []int {
	var pivot int
	var parIndex int

	if left < right {
		pivot = right
		parIndex = left

		for i := left; i < right; i++ {
			if arr[i] < arr[pivot] {
				swap(arr, i, parIndex)
				parIndex++
			}
		}
		swap(arr, parIndex, pivot)
		quickSort(arr, left, parIndex-1)
		quickSort(arr, parIndex+1, right)
	}
	return arr
}

// QuickSort implements the stated formula
func QuickSort(arr []int) []int {
	return quickSort(arr, 0, len(arr)-1)
}

func main() {
	a := time.Now()
	var arr [100]int
	var slice = arr[:]

	// set up random
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	// initialize an array with random numbers from 0 to 10000000
	for index := range arr {
		arr[index] = random.Intn(10000000)
	}

	// output before sort
	fmt.Printf("Before: %v\n", arr)

	b := time.Now()
	// sorting
	result := QuickSort(slice)

	// output after sort
	fmt.Printf("After: %v\n", result)
	fmt.Printf("Time for sort to complete: %v\n", time.Since(b))
	fmt.Printf("Time overall: %v\n", time.Since(a))

}
