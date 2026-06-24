package validanagram

func isAnagram(s string, t string) bool {

	if len(s) != len(t) {
		return false
	}

	seeMap := make(map[rune]int)

	for _, e := range s {
		seeMap[e]++
	}

	for _, e := range t {
		v, ok := seeMap[e]
		if !ok || v == 0 {
			return false
		}
		seeMap[e] = seeMap[e] - 1
	}

	return true
}

/*
Notes

Problem Link:
https://leetcode.com/problems/valid-anagram

*/
