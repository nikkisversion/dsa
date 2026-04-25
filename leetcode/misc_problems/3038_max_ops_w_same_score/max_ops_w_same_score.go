package maxopswsamescore

func maxOperations(nums []int) int {

	if len(nums) < 2 {
		return 0
	}

	res := 1

	for i := 2; i <= len(nums)-2; i = i + 2 {

		if nums[i]+nums[i+1] == nums[0]+nums[1] {
			res = res + 1
		} else {
			break
		}

	}

	return res

}

/*
Notes

Problem Link:
https://leetcode.com/problems/maximum-number-of-operations-with-the-same-score-i

Comments:
1. Take care of how to iterate over the loop. The starting and
limiting values of i, as well as the increment.

*/
