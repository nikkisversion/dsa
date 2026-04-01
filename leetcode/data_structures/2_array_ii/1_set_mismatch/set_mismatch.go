package setmismatch

func findErrorNums(nums []int) []int {

	freqMap := make(map[int]bool)

	n := len(nums)
	dup := 0
	actualSum := 0
	expectedSum := (n * (n + 1)) / 2

	for i := 0; i < n; i++ {
		if freqMap[nums[i]] {
			dup = nums[i]
		}
		freqMap[nums[i]] = true
		actualSum = actualSum + nums[i]
	}

	diff := expectedSum - actualSum
	missing := dup + diff

	return []int{dup, missing}

}

/*
Notes

-- Comments:

1. The set can be sorted/unsorted. So we need a
hashmap to keep a tarck of elements that we have
encountered. Since only one element can be
duplicated and that too only once, an int to bool
hashmap works well.

2. The sum of first n natural numbers is
n(n + 1) / 2. On trying out this question with
several example sets, we can conclude that the
different between actual and expected sum of
these n numbers will be equal to the difference
between the missing number and the duplicate.
Which means that
the missing number = (expected sum - actual sum) + duplicate.


-- Problem Link:
https://leetcode.com/problems/set-mismatch?envType=problem-list-v2&envId=dsa-linear-shoal-array-ii

*/
