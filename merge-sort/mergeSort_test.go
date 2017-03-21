package main

import (
	"reflect"
	"testing"
)

func TestMergesort(t *testing.T) {
	cases := []struct {
		in   []int
		want []int
	}{
		{[]int{5, 4, 3, 2, 1},
			[]int{1, 2, 3, 4, 5}},

		{[]int{6039, 5039, 4837, 3382, 100},
			[]int{100, 3382, 4837, 5039, 6039}},
	}
	for _, c := range cases {
		got := Mergesort(c.in)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("MergeSort(%v) == %v, want %v", c.in, got, c.want)
		}
	}
}
