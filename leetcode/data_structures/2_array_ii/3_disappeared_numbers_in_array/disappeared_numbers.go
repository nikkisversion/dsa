package disappearednumbersinarray

import (
	"fmt"
	"slices"
)

func findDisappearedNumbers(nums []int) []int {

	slices.Sort(nums)
	n := len(nums)
	res := []int{}

	fmt.Println(nums)
	counter := 1
	for i := 0; i < n; i++ {
		e := nums[i]
		fmt.Printf("i: %v, e: %v; counter: %v\n", i, e, counter)
		if e == counter {
			counter++
			continue
		}

		if e < counter {
			continue
		}

		//diff := e - counter
		for j := counter; j < e; j++ {
			res = append(res, j)
		}
		counter = e + 1
	}

	if nums[n-1] < n {
		for x := counter; x <= n; x++ {
			res = append(res, x)
		}
	}
	return res
}

func findDisappearedNumbers2(nums []int) []int {

	n := len(nums)
	seeMap := make(map[int]bool)

	for _, e := range nums {
		if seeMap[e] {
			continue
		}
		seeMap[e] = true
	}

	res := []int{}
	for i := 1; i < n+1; i++ {
		if !seeMap[i] {
			res = append(res, i)
		}
	}
	return res
}
