package maxfreqstack

import "container/list"

type FNode struct {
	Freq      int
	VNodeList *list.List
	FNodePtr  *list.Element
}

type VNode struct {
	Data     int
	FNodePtr *list.Element
	VNodePtr *list.Element
}

type FreqStack struct {
	FNodesList          *list.List
	ValueToFrequencyMap map[int]*FNode
}

func Constructor() FreqStack {
	return FreqStack{}
}

func (this *FreqStack) Push(val int) {

}

func (this *FreqStack) Pop() int {
	return 0
}
