package insertdeletegetrandomo1

import (
	"math/rand/v2"
)

type RandomizedSet struct {
	StorageMap   map[int]int
	StorageArray []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		StorageMap:   make(map[int]int),
		StorageArray: []int{},
	}
}

func (this *RandomizedSet) Insert(val int) bool {

	if _, ok := this.StorageMap[val]; ok {
		return false
	}

	l := len(this.StorageArray)
	this.StorageArray = append(this.StorageArray, val)
	this.StorageMap[val] = l

	return true
}

func (this *RandomizedSet) Remove(val int) bool {

	if _, ok := this.StorageMap[val]; !ok {
		return false
	}

	index := this.StorageMap[val]
	l := len(this.StorageArray)
	e := this.StorageArray[l-1]

	// swap the element to be deleted with the last
	this.StorageArray[index] = e

	// update position of last element
	this.StorageMap[e] = index

	delete(this.StorageMap, val)
	this.StorageArray = this.StorageArray[:l-1]

	return true
}

func (this *RandomizedSet) GetRandom() int {

	randomIndex := rand.IntN(len(this.StorageArray))
	return this.StorageArray[randomIndex]

}

/*
Notes

Problem Link:
https://leetcode.com/problems/insert-delete-getrandom-o1?envType=problem-list-v2&envId=ssd-ssd3-data-structure-design

Comments:

1. We are using a map to ensure Insert in O(1) time.

2. To have getRandom in O(1) time with all elements
having equal probability, we need an array.

3. To ensure removal takes O(1), we swap the element
to be removed with the last element. Then we pop the
last element from the array. It is important to also
update the index of the last element in the map. This
allows us to retain the order of elements in the array,
which means that we won't have to update each element's
index in the map after the removal of one element.

*/
