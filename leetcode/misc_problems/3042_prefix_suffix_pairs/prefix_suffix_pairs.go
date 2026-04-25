package prefixsuffixpairs

import "strings"

func countPrefixSuffixPairs(words []string) int {

	pairs := 0

	for i := 0; i < len(words); i++ {
		e := words[i]

		for j := i + 1; j < len(words); j++ {

			s := words[j]

			if isPrefixAndSuffix(e, s) {
				pairs++
			}
		}
	}

	return pairs

}

func isPrefixAndSuffix(str1, str2 string) bool {
	return strings.HasPrefix(str2, str1) && strings.HasSuffix(str2, str1)
}

/*
Notes

Problem Link:
https://leetcode.com/problems/count-prefix-and-suffix-pairs-i

Comments:
1. Check if there is a way to do it without nested loop.
Some solutions mention trie.

References:
Google 'check if string is prefix golang' for
alternatives to identify prefixes/suffixes that do not
use library packages.

*/
