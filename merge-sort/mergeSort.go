package main // In computer science, merge sort (also commonly spelled mergesort) is an efficient, general-purpose, comparison-based
import (
	"fmt"
	"math/rand"
	"time"
)

// sorting algorithm. Most implementations produce a stable sort, which means that the implementation preserves the
// input order of equal elements in the sorted output. Mergesort is a divide and conquer algorithm that was invented by
// John von Neumann in 1945.
// Conceptually, a merge sort works as follows:
// Divide the unsorted list into n sublists, each containing 1 element (a list of 1 element is considered sorted).
// Repeatedly merge sublists to produce new sorted sublists until there is only 1 sublist remaining. This will be the
// sorted list.
// In sorting n objects, merge sort has an average and worst-case performance of O(n log n). If the running time of
// merge sort for a list of length n is T(n), then the recurrence T(n) = 2T(n/2) + n follows from the definition of the
// algorithm (apply the algorithm to two lists of half the size of the original list, and add the n steps taken to
// merge the resulting two lists). The closed form follows from the master theorem.
// In the worst case, the number of comparisons merge sort makes is equal to or slightly smaller than (n ⌈lg n⌉ - 2⌈lg
// n⌉ + 1), which is between (n lg n - n + 1) and (n lg n + n + O(lg n)).
// In the worst case, merge sort does about 39% fewer comparisons than quicksort does in the average case. In terms of
// moves, merge sort's worst case complexity is O(n log n)—the same complexity as quicksort's best case, and merge
// sort's best case takes about half as many iterations as the worst case.
// Merge sort is more efficient than quicksort for some types of lists if the data to be sorted can only be efficiently
// accessed sequentially, and is thus popular in languages such as Lisp, where sequentially accessed data structures
// are very common. Unlike some (efficient) implementations of quicksort, merge sort is a stable sort.
// Merge sort's most common implementation does not sort in place;[5] therefore, the memory size of the input must be
// allocated for the sorted output to be stored in (see below for versions that need only n/2 extra spaces).
// source : https://en.wikipedia.org/wiki/Merge_sort

// MergeSort takes a slice and sorts them using the above algo
func Mergesort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	var mid int = len(arr) / 2
	var left []int
	var right []int

	// make copies for the left and right side
	left = make([]int, len(arr[:mid]))
	copy(left, Mergesort(arr[:mid])) // copy left side recursively

	right = make([]int, len(arr[mid:]))
	copy(right, Mergesort(arr[mid:])) // copy right side recursively

	var i int
	var j int
	var k int

	// merge using 2 pointers on the left and right to compare
	for (i < len(left)) && (j < len(right)) {
		if left[i] < right[j] {
			arr[k] = left[i]
			i++
		} else {
			arr[k] = right[j]
			j++
		}
		k++
	}
	// process the left block items
	for i < len(left) {
		arr[k] = left[i]
		i++
		k++
	}

	// process the right block items
	if j < len(right) {
		arr[k] = right[j]
		j++
		k++
	}
	return arr
}

func main() {
	a := time.Now()
	var arr [100]int
	var slice = arr[:]

	// set up rand numbers
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	// initialize array with random numbers from 0 to 10000000
	for index := range arr {
		arr[index] = random.Intn(10000000)
	}

	// output before sort
	fmt.Printf("Before: %v\n", arr)
	b := time.Now()
	// sorting
	result := Mergesort(slice)

	// output after
	fmt.Printf("After: %v\n", result)
	fmt.Printf("Time for sort to complete: %v\n", time.Since(b))
	fmt.Printf("Time overall: %v\n", time.Since(a))
}
