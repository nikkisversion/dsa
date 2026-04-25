package splitarray

func isPossibleToSplit(nums []int) bool {

	freqMap := make(map[int]int)

	for _, e := range nums {

		f, ok := freqMap[e]

		if !ok {
			freqMap[e] = 1
			continue
		}

		f = f + 1
		if f >= 3 {
			return false
		}

		freqMap[e] = f
	}
	return true
}

/*
Notes

Problem Link:
https://leetcode.com/problems/split-the-array

Comments:
1. If any element has a frequency of 3 or over,
then the array cannot be split with distinct
elements. Thus, we really only need one pass
of the array to find this out.

*/
