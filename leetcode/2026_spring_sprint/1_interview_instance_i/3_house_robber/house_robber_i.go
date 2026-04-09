package houserobber

// rob(i) = max(rob(i - 1), rob(i-2) + current)
var maxLoots []int

func rob(nums []int) int {

	maxLoots = make([]int, len(nums))
	for i, _ := range maxLoots {
		maxLoots[i] = -1
	}

	return getMaxLoot(nums, len(nums)-1)

}

func getMaxLoot(nums []int, i int) int {

	if i < 0 {
		return 0
	}

	if maxLoots[i] >= 0 {
		return maxLoots[i]
	}

	res1 := getMaxLoot(nums, i-2)
	res2 := getMaxLoot(nums, i-1)

	res1 = res1 + nums[i]

	res := max(res1, res2)
	maxLoots[i] = res

	return res

}

/*
Notes

Problem Link:
https://leetcode.com/problems/house-robber

References:
https://leetcode.com/problems/house-robber/solutions/156523/from-good-to-great-how-to-approach-most-ie2yi

*/
