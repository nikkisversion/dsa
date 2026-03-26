package main

import (
	"container/list"
	"fmt"
)

type Node struct {
	Data int
	Ptr  *list.Element
}

type LRUCache struct {
	Capacity int
	Queue    *list.List
	Items    map[int]*Node
}

func Constructor(capacity int) LRUCache {
	queue := list.New()
	items := make(map[int]*Node)

	return LRUCache{
		Capacity: capacity,
		Queue:    queue,
		Items:    items,
	}
}

func (this *LRUCache) Get(key int) int {

	item, ok := this.Items[key]

	if !ok {
		return -1
	}

	// item exists: move to front of queue and return data
	this.Queue.MoveToFront(item.Ptr)
	return item.Data
}

func (this *LRUCache) Put(key int, value int) {

	item, ok := this.Items[key]

	// item exists: update data and move item to front of list
	if ok {
		item.Data = value
		this.Items[key] = item
		this.Queue.MoveToFront(item.Ptr)
	} else {
		// item does not exist: check capacity of LRUCache and then add new item

		if this.Capacity == len(this.Items) {
			lastItem := this.Queue.Back()
			this.Queue.Remove(lastItem)
			delete(this.Items, lastItem.Value.(int))
		}

		newQueueElement := this.Queue.PushFront(key)
		this.Items[key] = &Node{Data: value, Ptr: newQueueElement}
	}

}

func main() {

	lru := Constructor(3)

	lru.Put(1, 1)
	lru.Put(2, 2)
	fmt.Println(lru.Get(1))
	lru.Put(3, 3)
	fmt.Println(lru.Get(3))
	lru.Put(4, 4)
	fmt.Println(lru.Get(2))
}

/* Notes:

-- Problem Link: https://leetcode.com/problems/lru-cache?envType=problem-list-v2&envId=ssd-ssd1-cache-system-design

-- Comments:
Notice how while adding a new element to the queue,
the field Value of the conatiner/list.Element struct
is used as the key of the pair. This comes in useful
when we have to remove the last element of the queue.
Because then we get the last item of the list, get
the key from that item, and use that key to delete
the item from the map as well.
*/
