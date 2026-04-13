package reversepolishnotation

import (
	"fmt"
	"strconv"
)

type Stack struct {
	storage []int
	length  int
}

func (this *Stack) Push(n int) {
	this.storage = append(this.storage, n)
	this.length = this.length + 1
}

func New() *Stack {
	return &Stack{
		storage: []int{},
		length:  0,
	}
}

func (this *Stack) Pop() int {
	e := this.storage[this.length-1]
	this.storage = this.storage[:this.length-1]
	this.length = this.length - 1
	return e
}

func evalRPN1(tokens []string) int {

	stack := New()
	//l := 0

	for _, e := range tokens {

		switch e {
		case "+":
			e1 := stack.Pop()
			e2 := stack.Pop()
			r := e1 + e2
			fmt.Printf(" + : e1 -> %v, e2 -> %v, r -> %v", e1, e2, r)
			stack.Push(r)
		case "-":
			e1 := stack.Pop()
			e2 := stack.Pop()
			r := e2 - e1
			stack.Push(r)
			fmt.Printf(" - : e1 -> %v, e2 -> %v, r -> %v", e1, e2, r)
		case "*":
			e1 := stack.Pop()
			e2 := stack.Pop()
			r := e1 * e2
			stack.Push(r)
			fmt.Printf(" * : e1 -> %v, e2 -> %v, r -> %v", e1, e2, r)
		case "/":
			e1 := stack.Pop()
			e2 := stack.Pop()
			r := e2 / e1
			stack.Push(r)
			fmt.Printf(" / : e1 -> %v, e2 -> %v, r -> %v", e1, e2, r)
		default:
			num, _ := strconv.Atoi(e)
			stack.Push(num)
		}

	}

	return stack.Pop()

}
