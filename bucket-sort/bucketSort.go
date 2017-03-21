// Bucket sort, or bin sort, is a sorting algorithm that works by distributing the elements of an array into a number
// of buckets. Each bucket is then sorted individually, either using a different sorting algorithm, or by recursively
// applying the bucket sorting algorithm. It is a distribution sort, and is a cousin of radix sort in the most to least
// significant digit flavor. Bucket sort is a generalization of pigeonhole sort. Bucket sort can be implemented with
// comparisons and therefore can also be considered a comparison sort algorithm. The computational complexity estimates
// involve the number of buckets.
// Bucket sort works as follows:
// Set up an array of initially empty "buckets".
// Scatter: Go over the original array, putting each object in its bucket.
// Sort each non-empty bucket.
// Gather: Visit the buckets in order and put all elements back into the original array.
// Note that for bucket sort to be {\displaystyle O(n)} O(n) on average, the number of buckets n must be equal to the
// length of the array being sorted, and the input array must be uniformly distributed across the range of possible
// bucket values. If these requirements are not met, the performance of bucket sort will be dominated by the running
// time of nextSort, which is typically {\displaystyle O(n^{2})} O(n^{2}) insertion sort, making bucket sort less
// optimal than {\displaystyle O(n\log(n))} O(n\log(n)) comparison sort algorithms like Quicksort. // Source ::
// https://en.wikipedia.org/wiki/Bucket_sort

package main

import (
	"fmt"
	"math/rand"
	"time"
)

// BucketSort takes a slice of items and sorts them using the above algo
func BucketSort(arr []int) []int {
	var largest int
	var bucket []int
	var result []int

	// Initialize result
	resultLength := len(arr)
	result = make([]int, resultLength, resultLength)
	// find the largest int
	fmt.Println(largest)
	for _, v := range arr {
		if v > largest {
			largest = v
			fmt.Println(largest)
		}
	}
	fmt.Println(largest)

	// initialize bucket
	bucket = make([]int, largest+1, largest+1)

	// add value to the bucket
	for _, v := range arr {
		bucket[v]++
	}

	// count the numbers inside the bucket and output to the result
	resultPointer := 0
	for i, v := range bucket {
		for j := 0; j < v; j++ {
			result[resultPointer] = i
			resultPointer++
		}
	}

	return result
}

func main() {
	a := time.Now()
	var arr [100]int
	var slice = arr[:]

	// set up rand number
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	// initialize array with a random number from 0 to 1000000000
	for index := range arr {
		arr[index] = random.Intn(1000000000)
	}
	// output before sort
	fmt.Printf("Before : %v\n", arr)
	// sorting
	b := time.Now()
	result := BucketSort(slice)
	// output after sort
	fmt.Printf("After: %v\n", result)
	fmt.Printf("Time for sort to complete: %v\n", time.Since(b))
	fmt.Printf("Time overall: %v\n", time.Since(a))
}
