package heightchecker

import "slices"

func heightChecker(heights []int) int {
	expected := []int{}
	expected = append(expected, heights...)
	slices.Sort(expected)

	res := 0

	for i, _ := range heights {
		if heights[i] != expected[i] {
			res++
		}
	}

	return res
}

/*
Notes

Problem Link:
https://leetcode.com/problems/height-checker

Comments:

1. The main thing here is to note how to form the
slice of expected heights. On simply assigning
the original slice to the new, we make a shallow
copy, which means it points to the same
underlying slice. This results in the original
being modified when the expected heights are
sorted. Hence we need an alternative. This can
be the methos used here. Or we can do
slices.Clone(), or the copy() method.

*/
