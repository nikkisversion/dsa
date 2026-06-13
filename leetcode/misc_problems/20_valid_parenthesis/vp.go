package validparenthesis

func isValid(s string) bool {
	pMap := make(map[rune]rune)
	pMap[')'] = '('
	pMap['}'] = '{'
	pMap[']'] = '['

	stack := []rune{}

	for _, char := range s {

		// if opening parenthesis, then push to stack
		if char == '{' || char == '[' || char == '(' {
			stack = append(stack, char)
			continue
		}

		v, ok := pMap[char]
		if !ok || len(stack) == 0 {
			// invalid character
			return false
		}

		e := stack[len(stack)-1]
		if v == e {
			stack = stack[:len(stack)-1]
			continue
		} else {
			return false
		}
	}

	if len(stack) > 0 {
		return false
	}

	return true
}

/*

Problem Link: https://leetcode.com/problems/valid-parentheses

*/
