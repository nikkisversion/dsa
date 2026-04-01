package maxconsecutive1s

func findMaxConsecutiveOnes(nums []int) int {

	m := 0
	maxi := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			m++
		} else {
			maxi = max(m, maxi)
			m = 0
		}
	}
	maxi = max(m, maxi)
	return maxi
}

/*

Notes

-- Comments:
1. A sliding window approach is suggested
2. This approach is not that of a traditional
sliding window in the sense that we are not
keeping a track of the left and right pointers.
But we are still maintaining a sum of a particular
window in the variable m.

-- Problem Link:
https://leetcode.com/problems/max-consecutive-ones?envType=problem-list-v2&envId=dsa-linear-shoal-array-i

*/
