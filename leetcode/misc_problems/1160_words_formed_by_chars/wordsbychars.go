package wordsformedbychars

import "slices"

func countCharacters(words []string, chars string) int {

	freqArray := make([]rune, 26)

	for _, e := range chars {
		freqArray[e-'a']++
	}

	res := 0
	lenChars := len(chars)
	for _, word := range words {

		l := len(word)
		if l > lenChars {
			continue
		}

		temp := slices.Clone(freqArray)
		if wordExistsInString(word, chars, temp) {
			res = res + l
		}
	}

	return res

}

func wordExistsInString(word, str string, freqArray []rune) bool {

	for _, e := range word {
		if freqArray[e-'a'] == 0 {
			return false
		}
		freqArray[e-'a']--
	}

	return true
}

/*
Notes

Problem Link:
https://leetcode.com/problems/find-words-that-can-be-formed-by-characters

Related Problem Link:
https://leetcode.com/problems/ransom-note

Comments:

1. Refer to #383 Ransom Note problem notes.

2. A catch is to use a copy of the frequency array
to decrement while comparing so as to not change
the original frequency array. This is important as
we are comparing several words and would need the
original array for each word.

3. Time complexity: O(m x n x l) where
m: length of chars
n: length of words array
l: length of each word

*/
