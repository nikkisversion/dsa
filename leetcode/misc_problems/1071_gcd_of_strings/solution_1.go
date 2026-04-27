package gcdofstrings

import (
	"fmt"
	"strings"
)

func gcdOfStrings1(str1 string, str2 string) string {

	l1 := len(str1)
	l2 := len(str2)

	bigger := ""
	smaller := ""

	if l1 <= l2 {
		bigger = str2
		smaller = str1
	} else {
		bigger = str1
		smaller = str2
	}

	gcd := ""
	strToCheck := ""
	for _, e := range smaller {

		strToCheck = strToCheck + string(e)
		if strings.HasPrefix(bigger, strToCheck) {
			if isDivisor(bigger, strToCheck) && isDivisor(smaller, strToCheck) {
				gcd = strToCheck
			}
		} else {
			return ""
		}
	}
	return gcd
}

func isDivisor(str, x string) bool {

	if !strings.HasPrefix(str, x) {
		return false
	}

	if len(str)%len(x) != 0 {
		return false
	}

	l := len(x)

	for i := 0; i <= len(str)-l; i = i + l {
		y := str[i : i+l]
		fmt.Printf("IsDivisor - y: %v, x: %v\n", y, x)
		if y != x {
			return false
		}
	}

	return true

}

/*
Notes

Problem Link:
https://leetcode.com/problems/greatest-common-divisor-of-strings

References:
https://leetcode.com/problems/greatest-common-divisor-of-strings/solutions/6601478/euclids-algorithm-visualized-step-by-ste-ftwo

*/
