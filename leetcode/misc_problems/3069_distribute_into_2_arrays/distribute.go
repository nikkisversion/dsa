package distributeinto2arrays

func resultArray(nums []int) []int {

	if len(nums) <= 2 {
		return nums
	}

	arr1 := []int{nums[0]}
	arr2 := []int{nums[1]}

	l1 := 1
	l2 := 1

	for i := 2; i < len(nums); i++ {

		e := nums[i]

		if arr1[l1-1] > arr2[l2-1] {
			arr1 = append(arr1, e)
			l1++
		} else {
			arr2 = append(arr2, e)
			l2++
		}

	}

	res := append(arr1, arr2...)
	return res
}

/*
Notes

Problem Link:
https://leetcode.com/problems/distribute-elements-into-two-arrays-i

*/
