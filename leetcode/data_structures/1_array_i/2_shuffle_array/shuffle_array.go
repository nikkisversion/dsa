package shufflearray

func shuffle(nums []int, n int) []int {
	a1 := nums[0:n]
	a2 := nums[n:]

	res := []int{}
	for i := 0; i < n; i++ {
		res = append(res, a1[i])
		res = append(res, a2[i])
	}
	return res
}

/*
Notes

-- Comments:

1. The first approach was to keep 2 pointers: one
at the start of the list and one at the nth element.
And then in a for loop, 2n times, to check the mod
of i and use the corresponding pointer to fill
the resulting array.

2. The other appraoch, which cut the execution time
in half, was to split the array into two and then
run the for loop n times and append an element
from both arrays to the new one.


-- Problem Link:
https://leetcode.com/problems/shuffle-the-array?envType=problem-list-v2&envId=dsa-linear-shoal-array-i

*/
