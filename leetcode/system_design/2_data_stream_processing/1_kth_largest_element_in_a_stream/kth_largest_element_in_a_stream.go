package kthlargestelementinastream

import "slices"

type KthLargest struct {
	K    int
	Nums []int
}

func Constructor(k int, nums []int) KthLargest {
	slices.Sort(nums)
	return KthLargest{K: k, Nums: nums}
}

func (this *KthLargest) Add(val int) int {

	index, _ := slices.BinarySearch(this.Nums, val)
	this.Nums = slices.Insert(this.Nums, index, val)

	req := len(this.Nums) - this.K
	return this.Nums[req]
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */

/* Notes

-- Problem Link:
https://leetcode.com/problems/kth-largest-element-in-a-stream?envType=problem-list-v2&envId=ssd-ssd2-data-stream-processing

-- Comments:
Keep in mind how the index of the kth largest element is calculated.
*/
