package main

import "fmt"

func main() {
	result := distributeCandies(10, 7)
	fmt.Println(result)
}

func distributeCandies(candies int, num_people int) []int {
	result := make([]int, num_people)

	index := 0
	for i := 1; candies > 0; i++ {

		result[index] += min(i, candies)
		candies = candies - i

		index = index + 1
		if index >= num_people {
			index = index % num_people
		}
	}
	return result
}
