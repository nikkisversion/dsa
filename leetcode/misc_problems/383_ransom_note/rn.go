package ransomnote

func canConstruct(ransomNote string, magazine string) bool {

	freqArray := make([]rune, 26)

	for _, e := range magazine {
		freqArray[e-'a']++
	}

	for _, e := range ransomNote {
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
https://leetcode.com/problems/ransom-note

Comments:

1. One way would be to make a frequency map of both
strings and then to compare the frequency of each
character in the ransom string. If the frequency of
this character in the magazine string is greater
than or equal to the ransom one, then we can
proceed. Else, it won't be possible to make the
ransom string from the magazine one.

But this is of
a bigger time complexity as we will be traversing
through both strings once for creating the frequency
maps and then again through one map to compare.

We
can shorten this by comparing as we build the
frequency map of the ransom string, but it is still
time complexity O(m x n), where m and n would be
the sizes of the strings.


2. A more straightforward approach is to create a
fixed-length array to store the frequencies of
each character in the magazine string, and keep
decrementing this as we move throught the ransom
string. O(m x n) again, but takes lesser time
in leetcode.


3. Note how we are calculating the index of the
rune in the string, by subtracting 'a' from it.
This gives us a numerical value starting from
0 for 'a'.

4. We can also use strings.Count() library
function.

*/
