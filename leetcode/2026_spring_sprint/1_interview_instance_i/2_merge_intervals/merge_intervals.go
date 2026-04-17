package mergeintervals

import (
	"cmp"
	"slices"
)

func sortFunc(a, b []int) int {
	if n := cmp.Compare(a[0], b[0]); n != 0 {
		return n
	}
	return cmp.Compare(a[1], b[1])
}

func canMerge(a, b []int) bool {

	// sorted so a will be smaller than e

	if a[0] == b[0] {
		return true
	}

	if b[0] <= a[1] {
		return true
	}

	return false
}

func formMerge(a, b []int) []int {

	res := []int{a[0]}

	if a[1] > b[1] {
		res = append(res, a[1])
	} else {
		res = append(res, b[1])
	}

	return res

}

func merge(intervals [][]int) [][]int {

	slices.SortFunc(intervals, sortFunc)
	//fmt.Println(intervals)

	res := [][]int{intervals[0]}
	j := 0

	for i := 1; i < len(intervals); i++ {

		// sorted so target will be smaller than e
		e := intervals[i]
		target := res[j]

		if target[0] == e[0] || e[0] <= target[1] {
			//merge
			//new := formMerge(target, e)
			new := []int{target[0], max(target[1], e[1])}
			//res = res[:len(res) - 1]
			//res = append(res, new)
			res[j] = new
		} else {
			res = append(res, e)
			j++
		}
	}

	return res
}

/*
Notes

Problem link:
https://leetcode.com/problems/merge-intervals?envType=problem-list-v2&envId=interview-instance-i

Comments:

1. The method is to first have the array sorted.
Once sorted, you can traverse the array by element,
and decide whether to merge it or not.

2. Since we are sorting before traversing, this
allows us to know that when we have an element
we are deciding to merge or not, we will have to
check that element with only the last element of
our result array.

*/
