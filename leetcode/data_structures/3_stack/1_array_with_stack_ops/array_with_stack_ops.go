package arraywithstackops

func buildArray(target []int, n int) []string {

	var result []string

	lenTarget := len(target)
	j := 0 // iterator for target array

	for i := 0; i < n; i++ {

		if j == lenTarget {
			break
		}

		if i+1 == target[j] {
			result = append(result, "Push")
			j++
		} else {
			result = append(result, "Push", "Pop")
		}

	}

	return result
}

/*
Notes

Problem Link:
https://leetcode.com/problems/build-an-array-with-stack-operations?envType=problem-list-v2&envId=dsa-linear-shoal-stack

Notes:

1. The target array is in increasing order.
It is mentioned in the question.

2. This is not a stack problem. It is a problem
where you compare a given array (an array of
numbers from 1 to n, representing the 'stream')
to a target array. Your goal is to identify
which 'stack' operations will take place to
convert the given array to target.

3. You go over the given n numbers
(you can also make an array of numbers from
1 to n and then go over that, I refined the
approach to go over numbers from 1 to n only).

4. For each number, if it exists in the target
array, you append 'Push' to the result and
increment the target array counter 'j'. Else,
you append 'Push' and 'Pop' to the result,
and don't increment 'j', as we still have to
find the target element (target[j]).

*/
