package largestsubstringbw2equalchars

func maxLengthBetweenEqualCharacters(s string) int {
	seeMap := make(map[string][]int)

	for i, char := range s {

		e := string(char)

		val, ok := seeMap[e]
		if !ok {
			seeMap[e] = []int{i}
		} else {
			seeMap[e] = append(val, i)
		}

	}

	maxLen := -1
	for _, indexes := range seeMap {
		l := len(indexes)

		if l == 0 {
			continue
		}

		maxLenChar := indexes[l-1] - indexes[0] - 1

		if maxLenChar >= maxLen {
			maxLen = maxLenChar
		}
	}

	return maxLen
}

/*
Notes

Problem Link:
https://leetcode.com/problems/largest-substring-between-two-equal-characters

Comments:

1. We only need to keep a track of the indices of every pair
of elements that are equal.

2. If there are more than 2 occurrences of the same element
in the string, then the length of max substring will be the
difference between its first and last occurrence.

*/
