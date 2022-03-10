package main

import "strconv"

func IsPalindrome(number int) bool {
	if number < 0 {
		return false
	}

	numberStr := strconv.Itoa(number)

	lastIndex := len(numberStr) / 2

	for i := range numberStr[:lastIndex] {
		if numberStr[i] != numberStr[len(numberStr)-i-1] {
			return false
		}
	}

	return true
}
