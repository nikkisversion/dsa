package insertdeletegetrandomduplicates

import "math/rand/v2"

type RandomizedCollection struct {
	StorageArray []int
	StorageMap   map[int][]int
}

type Element struct {
	Indexes []int
}

func Constructor() RandomizedCollection {
	return RandomizedCollection{
		StorageArray: []int{},
		StorageMap:   make(map[int][]int),
	}
}

func (this *RandomizedCollection) Insert(val int) bool {

	_, ok := this.StorageMap[val]
	if ok {
		this.StorageArray = append(this.StorageArray, val)
		this.StorageMap[val] = append(this.StorageMap[val], len(this.StorageArray)-1)
		return false
	}

	this.StorageArray = append(this.StorageArray, val)
	this.StorageMap[val] = []int{len(this.StorageArray) - 1}
	return true
}

func (this *RandomizedCollection) Remove(val int) bool {

	indexes, ok := this.StorageMap[val]
	if !ok {
		return false
	}

	indexToRemove := indexes[len(indexes)-1]

	l := len(this.StorageArray)

	if l == 1 {
		delete(this.StorageMap, val)
		this.StorageArray = []int{}
		return true
	}

	if indexToRemove == l-1 {
		this.StorageArray = this.StorageArray[:l-1]
		newArray := indexes[:len(indexes)-1]
		this.StorageMap[val] = newArray
		return true
	}

	lastElementOfArray := this.StorageArray[l-1]

	this.StorageArray[indexToRemove] = lastElementOfArray
	this.StorageArray = this.StorageArray[:l-1]

	newIndexesArray := indexes[:len(indexes)-1]
	this.StorageMap[val] = newIndexesArray

	indexesForLastElement := this.StorageMap[lastElementOfArray]
	newIndexesForLastElement := indexesForLastElement[:len(indexesForLastElement)-1]
	newIndexesForLastElement = append(newIndexesForLastElement, indexToRemove)

	return true
}

func (this *RandomizedCollection) GetRandom() int {
	randomIndex := rand.IntN(len(this.StorageArray))
	return this.StorageArray[randomIndex]
}
