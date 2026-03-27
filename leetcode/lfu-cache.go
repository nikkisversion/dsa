package leetcode

import "container/list"

type ValueNode struct {
	Data          int
	Frequency     int
	FrequencyNode *list.Element
}

type FrequencyNode struct {
	Frequency int
	ValueList *list.List
}

type LFUCache struct {
	Capacity      int
	FrequencyList *list.List
	ItemsMap      map[int]*list.Element
	FrequencyMap  map[int]*list.Element
}

func Constructor(capacity int) LFUCache {
	frequencyList := list.New()
	itemsMap := make(map[int]*list.Element)
	freqMap := make(map[int]*list.Element)

	return LFUCache{
		Capacity:      capacity,
		FrequencyList: frequencyList,
		ItemsMap:      itemsMap,
		FrequencyMap:  freqMap,
	}
}

func (this *LFUCache) Get(key int) int {

	item, ok := this.ItemsMap[key]

	if !ok {
		return -1
	}

	valueNode := item.Value.(ValueNode)
	requiredValue := valueNode.Data

	this.updateFrequency(item)

	// currentFreq := valueNode.Frequency
	// newFreq := currentFreq + 1

	// currentFreqNode := this.FrequencyMap[currentFreq].Value.(FrequencyNode)
	// currentFreqList := currentFreqNode.ValueList
	// currentFreqList.Remove(item)

	// newFreqNode := this.FrequencyMap[newFreq].Value.(FrequencyNode)
	// newFreqList := newFreqNode.ValueList

	// newValueNode := ValueNode{
	// 	Data:          valueNode.Data,
	// 	Frequency:     newFreq,
	// 	FrequencyNode: this.FrequencyMap[newFreq],
	// }

	// newFreqList.PushFront(newValueNode)

	return requiredValue

}

func (this *LFUCache) Put(key int, value int) {

	item, ok := this.ItemsMap[key]

	// item exists, update value and frequency
	if ok {
		valueNode := item.Value.(ValueNode)
		valueNode.Data = value
		item.Value = valueNode
		this.updateFrequency(item)
	} else {
		newValueNode := ValueNode{Data: value, Frequency: 1}

		singleFreqNode, ok := this.FrequencyMap[1]
		if !ok {
			newFreqNode := FrequencyNode{
				Frequency: 1,
				ValueList: list.New(),
			}

			newFreqListElement := this.FrequencyList.PushFront(newFreqNode)
			newValueNode.FrequencyNode = newFreqListElement

			newValueNodeElement := newFreqNode.ValueList.PushFront(newValueNode)
			this.ItemsMap[key] = newValueNodeElement
		} else {
			newValueNode.FrequencyNode = singleFreqNode
			singleFreqNodeValueList := singleFreqNode.Value.(FrequencyNode).ValueList
			newValueNodeElement := singleFreqNodeValueList.PushFront(newValueNode)
			this.ItemsMap[key] = newValueNodeElement
		}
	}

}

func (this *LFUCache) updateFrequency(item *list.Element) {

	valueNode := item.Value.(ValueNode)
	currentFreq := valueNode.Frequency
	newFreq := currentFreq + 1

	currentFreqNode := this.FrequencyMap[currentFreq].Value.(FrequencyNode)
	currentFreqList := currentFreqNode.ValueList
	currentFreqList.Remove(item)

	newFreqNode := this.FrequencyMap[newFreq].Value.(FrequencyNode)
	newFreqList := newFreqNode.ValueList

	newValueNode := ValueNode{
		Data:          valueNode.Data,
		Frequency:     newFreq,
		FrequencyNode: this.FrequencyMap[newFreq],
	}

	newFreqList.PushFront(newValueNode)
}
