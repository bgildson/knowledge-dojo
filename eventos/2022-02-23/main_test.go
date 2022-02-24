package main

import (
	"reflect"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	tests := []struct {
		name string
		in1  []int
		in2  []int
		out  []int
	}{
		{
			name: "basic",
			in1:  []int{1, 2, 4},
			in2:  []int{1, 3, 4},
			out:  []int{1, 1, 2, 3, 4, 4},
		},
		{
			name: "empty inputs",
			in1:  []int{},
			in2:  []int{},
			out:  []int{},
		},
		{
			name: "one empty and one populated",
			in1:  []int{},
			in2:  []int{0},
			out:  []int{0},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := MergeTwoLists(test.in1, test.in2)
			if !reflect.DeepEqual(got, test.out) {
				t.Error("want: ", test.out, " got: ", got)
			}
		})
	}
}

func BenchmarkMergeTwoLists(b *testing.B) {
	tests := []struct {
		in1 []int
		in2 []int
	}{
		{
			in1: []int{1, 2, 4},
			in2: []int{1, 3, 4},
		},
		{
			in1: []int{},
			in2: []int{},
		},
		{
			in1: []int{},
			in2: []int{0},
		},
	}
	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			MergeTwoLists(test.in1, test.in2)
		}
	}
}
