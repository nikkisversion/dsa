package reversepolishnotation

import "strconv"

func evalRPN(tokens []string) int {

	stack := []int{}
	l := 0

	for _, e := range tokens {

		switch e {
		case "+":
			e1 := stack[l-1]
			e2 := stack[l-2]
			r := e1 + e2
			stack = stack[:l-2]
			stack = append(stack, r)
			l = l - 1
		case "*":
			e1 := stack[l-1]
			e2 := stack[l-2]
			r := e1 * e2
			stack = stack[:l-2]
			stack = append(stack, r)
			l = l - 1
		case "-":
			e1 := stack[l-1]
			e2 := stack[l-2]
			r := e2 - e1
			stack = stack[:l-2]
			stack = append(stack, r)
			l = l - 1
		case "/":
			e1 := stack[l-1]
			e2 := stack[l-2]
			r := e2 / e1
			stack = stack[:l-2]
			stack = append(stack, r)
			l = l - 1
		default:
			num, _ := strconv.Atoi(e)
			stack = append(stack, num)
			l++
		}

	}

	return stack[0]

}

/*
Notes

Problem Link:
https://leetcode.com/problems/evaluate-reverse-polish-notation?envType=problem-list-v2&envId=dsa-linear-shoal-stack

Comments:

1. We can use a proper Stack struct here,
like in the initial solution. But it
increases time complexity.

*/
