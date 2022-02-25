package main

import (
	"sort"
)

func Soma(x, y int) int {
	return x + y
}

func Fatorial(n int64) int64 {
	if n < 2 {
		return 1
	}
	return n * Fatorial(n-1)
}
func SearchBinaryIndex(sli []int, n int) int {
	sort.Ints(sli)
	start := 0
	end := len(sli)
	middle := (end - start) / 2
	for start < end {
        if sli[middle] < n {
            start = middle
			middle = middle + (end-start)/2
		} else {
			end = middle
            middle = (end - start) / 2
		}
		if sli[middle] == n {
			return middle
		}
        if middle == len(sli)-1 {
            return -1
        }
	}
	return -1
}
