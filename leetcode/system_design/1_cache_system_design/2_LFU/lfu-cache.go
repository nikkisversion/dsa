package lfu

import "container/list"

type VNode struct {
	Data     int
	Freq     int
	VNodePtr *list.Element
	FNodePtr *list.Element
}

type FNode struct {
	Freq      int
	FNodePtr  *list.Element
	VNodeList *list.List
}

type LFUCache struct {
	FNodeList *list.List
	Capacity  int
	ItemsMap  map[int]*VNode
	FreqMap   map[int]*FNode
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		FNodeList: list.New(),
		Capacity:  capacity,
		ItemsMap:  make(map[int]*VNode),
		FreqMap:   make(map[int]*FNode),
	}
}

func (this *LFUCache) Get(key int) int {

	vnode, ok := this.ItemsMap[key]

	if !ok {
		return -1
	}

	this.updateFrequency(key)
	return vnode.Data

}

func (this *LFUCache) Put(key int, value int) {

	_, ok := this.ItemsMap[key]

	if ok {
		this.updateFrequency(key)
		ref := this.ItemsMap[key]

		newVNode := VNode{
			Data:     value,
			Freq:     ref.Freq,
			VNodePtr: ref.VNodePtr,
			FNodePtr: ref.FNodePtr,
		}
		this.ItemsMap[key] = &newVNode
	} else {
		if this.Capacity == len(this.ItemsMap) {
			// evict one item
			// get least fnode = first element of fnodelist of lfucache
			// get least recently used vnode of this fnode = last element of its vnodelist
			// remove this vnode from this vnodelist and items map
			// if vnodelist is empty, remove this fnode from fnodelist and freqmap

			leastFreqNode := this.FNodeList.Front()
			leastFreq := leastFreqNode.Value.(int)
			leastFreqVNodeList := this.FreqMap[leastFreq].VNodeList

			lruVNode := leastFreqVNodeList.Back()
			lruKey := lruVNode.Value.(int)

			leastFreqVNodeList.Remove(lruVNode)
			delete(this.ItemsMap, lruKey)

			if leastFreqVNodeList.Len() == 0 {
				this.FNodeList.Remove(leastFreqNode)
				delete(this.FreqMap, leastFreq)
			}
		}

		singleFNode, ok := this.FreqMap[1]
		if ok {
			singleFNodeVList := singleFNode.VNodeList
			newVNodePtr := singleFNodeVList.PushFront(key)

			newVNode := VNode{
				Data:     value,
				Freq:     1,
				VNodePtr: newVNodePtr,
				FNodePtr: singleFNode.FNodePtr,
			}

			this.ItemsMap[key] = &newVNode
		} else {
			newFNodePtr := this.FNodeList.PushFront(1)

			newFNode := FNode{
				Freq:      1,
				FNodePtr:  newFNodePtr,
				VNodeList: list.New(),
			}

			newVNodePtr := newFNode.VNodeList.PushFront(key)

			newVNode := VNode{
				Data:     value,
				Freq:     1,
				FNodePtr: newFNodePtr,
				VNodePtr: newVNodePtr,
			}

			this.ItemsMap[key] = &newVNode
			this.FreqMap[1] = &newFNode
		}
	}

}

func (this *LFUCache) updateFrequency(key int) {

	vnode, ok := this.ItemsMap[key]
	if !ok {
		return
	}

	currentFreq := vnode.Freq
	newFreq := currentFreq + 1

	currentFNodePtr := vnode.FNodePtr
	currentVNodePtr := vnode.VNodePtr

	nextAvailableFreqNode := vnode.FNodePtr.Next()
	if nextAvailableFreqNode == nil {
		newFNodePtr := this.FNodeList.PushBack(newFreq)

		newFNode := FNode{
			Freq:      newFreq,
			FNodePtr:  newFNodePtr,
			VNodeList: list.New(),
		}

		newVNodePtr := newFNode.VNodeList.PushFront(key)

		newVnode := VNode{
			Data:     vnode.Data,
			Freq:     newFreq,
			VNodePtr: newVNodePtr,
			FNodePtr: newFNodePtr,
		}

		this.FreqMap[newFreq] = &newFNode
		this.ItemsMap[key] = &newVnode
	} else {
		// add new vnode to new freq
		nextAvailableFreq := nextAvailableFreqNode.Value.(int)
		if nextAvailableFreq == newFreq {
			// new freq node available. add new vnode to its vnodelist
			vnode.FNodePtr = this.FreqMap[newFreq].FNodePtr

			newVNodeList := this.FreqMap[newFreq].VNodeList
			newVNodePtr := newVNodeList.PushFront(key)

			vnode.VNodePtr = newVNodePtr
			vnode.Freq = newFreq

			// update items map with new vnode
			this.ItemsMap[key] = vnode
		} else {
			// new frequency node not available, create one
			newFNode := FNode{Freq: newFreq, VNodeList: list.New()}

			newFNodePtr := this.FNodeList.InsertAfter(newFreq, currentFNodePtr)
			newFNode.FNodePtr = newFNodePtr

			newVNodePtr := newFNode.VNodeList.PushFront(key)
			vnode.VNodePtr = newVNodePtr
			vnode.FNodePtr = newFNodePtr
			vnode.Freq = newFreq

			this.FreqMap[newFreq] = &newFNode
			this.ItemsMap[key] = vnode
		}
	}

	currentVNodeList := this.FreqMap[currentFreq].VNodeList
	currentVNodeList.Remove(currentVNodePtr)
	if currentVNodeList.Len() == 0 {
		// remove corresponding fnode from fnodelist and from freqmap of cache
		this.FNodeList.Remove(currentFNodePtr)
		delete(this.FreqMap, currentFreq)
	}

}

/* Notes:

-- Problem Link: https://leetcode.com/problems/lfu-cache?envType=problem-list-v2&envId=ssd-ssd1-cache-system-design

-- Comments:

1. When making FNode and VNode, we use FNodePtr and
VNodePtr. Here we are decoupling the data stored in
an FNode or VNode from the *list.Element type. This
is to make accessing these values easier. If we were
storing the whole struct as the Value field of each
*list.Element, then we would have had to perform
type conversion every time we wanted to access any
field in the VNode or FNode struct. This is also
parallel to how in LRU Cache we made Node struct
with Data and Ptr fields, to easily access Data
without accessing Ptr.

2. Take care of conditions where the element of the
list being access can be nil, like in updateFrequency
function while finding the next available frequency node.


3. Notice how while adding a new element to the queue,
the field Value of the conatiner/list.Element struct
is used as the key of the pair. This comes in useful
when we have to remove the last element of the queue.
Because then we get the last item of the list, get
the key from that item, and use that key to delete
the item from the map as well.

--- Useful Links:
	https://arpitbhayani.me/blogs/lfu/

*/
