package twosum

func twoSum(nums []int, target int) []int {
	seeMap := make(map[int]int)
	for i, e := range nums {
		newTarget := target - e
		if v, ok := seeMap[newTarget]; ok {
			return []int{v, i}
		}
		seeMap[e] = i
	}
	return []int{}
}

/*
Notes

Problem Link:
https://leetcode.com/problems/two-sum

TC: O(n) (for loop) + O(1) (map lookup) = O(n)
SC: O(n) for worst case (map)

*/
