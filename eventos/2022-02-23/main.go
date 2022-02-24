package main

import "sort"

func MergeTwoLists(list1 []int, list2 []int) []int {
	res := append(list1, list2...)
	sort.Ints(res)
	return res
}
