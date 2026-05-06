package gcdofstrings

func gcdOfStrings(str1 string, str2 string) string {

	if str1+str2 != str2+str1 {
		return ""
	}

	gcdOfLengths := gcd(len(str1), len(str2))

	return str1[:gcdOfLengths]

}

func gcd(a, b int) int {
	smaller := min(a, b)
	bigger := max(a, b)

	if smaller == 0 {
		return bigger
	} else {
		return gcd(smaller, bigger%smaller)
	}
}

/*
Notes

Problem Link:
https://leetcode.com/problems/greatest-common-divisor-of-strings

References:
https://leetcode.com/problems/greatest-common-divisor-of-strings/solutions/6601478/euclids-algorithm-visualized-step-by-ste-ftwo

*/
