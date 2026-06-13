package searchinrotatedsortedarray

func search(nums []int, target int) int {

	left := 0
	right := len(nums) - 1

	for left <= right {

		mid := (left + right) / 2

		if nums[mid] == target {
			return mid
		}

		if nums[left] <= nums[mid] {
			// left part is sorted
			if target >= nums[left] && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			// right part is sorted
			if target > nums[mid] && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}

/*
Notes

Problem Link: https://leetcode.com/problems/search-in-rotated-sorted-array

Comments:

1. Since the required time complexity is O(log n), this means we need to
use binary version in some form, even if modified.

2. Since binary search is basically halving the search bucket at every
step, we will use a modified version of it for the rotated sorted array.

3. Since the array is left rotated (or rotated in any direction), it
means that atleast one half of the array will be strictly sorted.

3. If the middle element is the target, return the index.

4. If not, then we see if the left part is sorted. If the left part
is sorted and target is within the left part, then we carry the search
within this part. Else in the right part.

5. If the left part is not sorted, this means the right one is. Now we
check if target is within this right half. If so, then we continue the
search in the right half, else we do it in the left half.

6. We return -1 if we exit the loop, which means target was not found.

*/
