package minopstoexceedthreshold

func minOperations(nums []int, k int) int {

	res := 0

	for _, e := range nums {
		if e < k {
			res++
		}
	}

	return res

}

/*
Notes

Problem Link:
https://leetcode.com/problems/minimum-operations-to-exceed-threshold-value-i

Comments:

1. Based on the description, we need to find out
the count of elements in the array that are
smaller than k.

*/
