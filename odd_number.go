package main

import (
	"sort"
)

// * Show coding logic and unit test Find the odd int (Choose your the most skilled)
// * Given an array of integers, find the one that appears an odd number of times.

// * There will always be only one integer that appears an odd number of times.

func FindOddNumber(text []int) int {
	// TODO : start your code here

	// SORT BEFORE FINDING ODD OCCUR NUMBER
	sort.Ints(text)
	result := -1
	currentNumberOccurs := 1
	if len(text) == 1 {
		return text[0]
	}
	for i, v := range text {
		if i == len(text)-1 && IsOdd(currentNumberOccurs) {
			// CHECK ODD TIME AND SET LAST NUMBER AS RESULT
			return v
		} else {
			// IS CHANGING NUMBER
			if v != text[i+1] {
				// CHECK ODD TIME AND SET LAST NUMBER AS RESULT
				if IsOdd(currentNumberOccurs) {
					// SET CURRENT NUMBER AS RESULT
					return v
				} else {
					// RESET OCCUR TIME
					currentNumberOccurs = 1
				}
			} else {
				currentNumberOccurs++
			}
		}
	}
	return result
}

func IsOdd(number int) bool {
	return number%2 != 0
}
