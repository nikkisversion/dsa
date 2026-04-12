package numberssmallerthaneach

import "slices"

func smallerNumbersThanCurrent(nums []int) []int {

	l := len(nums)
	positions := make([]int, l)

	for i, e := range nums {
		positions[i] = e
	}

	slices.Sort(nums)

	resMap := make(map[int]int)

	for i, e := range nums {
		_, ok := resMap[e]
		if !ok {
			resMap[e] = i
		}
	}

	result := make([]int, l)

	for i, _ := range result {
		element := positions[i]
		smallerThanElement := resMap[element]
		result[i] = smallerThanElement
	}

	return result

}

/*
Notes

Problem Link:
https://leetcode.com/problems/how-many-numbers-are-smaller-than-the-current-number?envType=problem-list-v2&envId=dsa-linear-shoal-array-ii

Notes:

1. One thing to note here is that if the array is sorted,
then the number of elements smaller than the current number
will be equal to its index. Or more correctly, the index
of its first occurrence (as there can be duplicates but
the number of elements smaller than, say, '2', will be
the same whether that 2 is at the index 3 or 4)

2. Upon sorting, we will lose the order of the input array.
This order is important to build the result. So we need
to preserve this order somewhere. I used a map based
approach at first (index -> element, to preserve order)
but that took more time than this approach of duplicating
the original array.

*/
