package concatenationofarray

func getConcatenation(nums []int) []int {

	n := len(nums)
	ans := make([]int, 2*n)

	for i := 0; i < n; i++ {
		ans[i] = nums[i]
		ans[i+n] = nums[i]
	}

	return ans
}

/*
Notes

-- Problem Link:
https://leetcode.com/problems/concatenation-of-array?envType=problem-list-v2&envId=dsa-linear-shoal-array-i

*/
