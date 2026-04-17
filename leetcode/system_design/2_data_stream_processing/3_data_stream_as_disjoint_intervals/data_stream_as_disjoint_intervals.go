package datastreamasdisjointintervals

import "slices"

type SummaryRanges struct {
	Storage []int
}

func Constructor() SummaryRanges {
	return SummaryRanges{Storage: []int{}}
}

func (this *SummaryRanges) AddNum(value int) {
	index, _ := slices.BinarySearch(this.Storage, value)
	this.Storage = slices.Insert(this.Storage, index, value)
}

func (this *SummaryRanges) GetIntervals() [][]int {

	if len(this.Storage) == 0 {
		return [][]int{}
	}

	result := [][]int{}

	l := 0
	r := 0
	for i := 0; i < len(this.Storage)-1; i++ {

		e := this.Storage[i]
		next := this.Storage[i+1]

		if next == e || next == e+1 {
			r++
			continue
		} else {
			result = append(result, []int{this.Storage[l], this.Storage[r]})
			l = i + 1
			r = i + 1
		}

	}
	result = append(result, []int{this.Storage[l], this.Storage[r]})
	return result
}

/*
Notes

Problem Link:
https://leetcode.com/problems/data-stream-as-disjoint-intervals?envType=problem-list-v2&envId=ssd-ssd2-data-stream-processing

Comments:

1. Adding is straightforward - keep the list sorted and add via binary search.

2. For counting intervals: do not store previous intervals,
else we will have to match incoming number with existing
intervals and update it on every addition which will increase
time complexity of add function and is also complex.

3. Instead, form the intervals when required. (Can compare
this to how the frequency of reads and writes might vary etc).
Go over the array in one pass maintaining two pointers for
the interval. When an interval's end is found, update the
pointers and append that interval. Take care of duplicates
by checking if the next element is equal to current or
current + 1. This will be complexity n as we are just
traversing the array once. Yay, such a good job!!

*/
