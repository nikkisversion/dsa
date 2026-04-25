package appleredistribution

import "slices"

func minimumBoxes(apple []int, capacity []int) int {

	totalApples := 0
	for _, e := range apple {
		totalApples += e
	}

	slices.Sort(capacity)

	cap := 0
	count := 0
	for i := len(capacity) - 1; i >= 0; i-- {
		e := capacity[i]

		if cap+e >= totalApples {
			count++
			return count
		}

		cap = cap + e
		count++
	}

	return count

}

/*
Notes

Problem Link:
https://leetcode.com/problems/apple-redistribution-into-boxes

Comments:

1. Since we need to find the minimum number fo boxes,
we can start by using the boxes of biggest capacity,
i.e., from the end of the sorted capacity array.

2. Once we reach a capacity that is >= the total
number of apples, we can stop.

3. The question description notes mention that all
the apples will be sorted. So we do not need to
account for a case where the number of apples are
greater than the available capacity.

*/
