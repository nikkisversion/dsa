package existenceofsubstrinstrandrev

import (
	"slices"
	"strings"
)

func isSubstringPresent(s string) bool {

	revStr := reverseString(s)

	l := len(s)

	for i := 0; i < l-1; i++ {
		x := string(s[i]) + string(s[i+1])
		if strings.Contains(revStr, x) {
			return true
		}
	}

	return false

}

func reverseString(s string) string {

	str := strings.Split(s, "")
	slices.Reverse(str)

	res := ""
	for _, e := range str {
		res = res + e
	}

	return res

}

/*
Notes

Problem Link:
https://leetcode.com/problems/existence-of-a-substring-in-a-string-and-its-reverse

*/
