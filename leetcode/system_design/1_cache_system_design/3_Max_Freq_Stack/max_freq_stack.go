package maxfreqstack

type FreqStack struct {
	FrequencyMap   map[int]int   // element to its frequency
	FrequencyStack map[int][]int // frequency to its elements, in order of push
	MaxFreq        int
}

func Constructor() FreqStack {
	return FreqStack{
		FrequencyMap:   make(map[int]int),
		FrequencyStack: make(map[int][]int),
		MaxFreq:        0,
	}
}

func (this *FreqStack) Push(val int) {

	currentFreq, ok := this.FrequencyMap[val]
	if !ok {
		this.FrequencyMap[val] = 1

		_, ok := this.FrequencyStack[1]
		if ok {
			this.FrequencyStack[1] = append(this.FrequencyStack[1], val)
		} else {
			this.FrequencyStack[1] = []int{val}
		}
	} else {
		// element already exists
		newFreq := currentFreq + 1

		_, ok := this.FrequencyStack[newFreq]
		if ok {
			this.FrequencyStack[newFreq] = append(this.FrequencyStack[newFreq], val)
		} else {
			this.FrequencyStack[newFreq] = []int{val}
		}

		this.FrequencyMap[val] = newFreq
	}

	if this.FrequencyMap[val] >= this.MaxFreq {
		this.MaxFreq = this.FrequencyMap[val]
	}

}

func (this *FreqStack) Pop() int {

	if len(this.FrequencyStack) == 0 {
		return 0
	}

	maxFreqStack, ok := this.FrequencyStack[this.MaxFreq]
	if !ok {
		return 0
	}

	l := len(maxFreqStack)
	e := maxFreqStack[l-1]

	if l == 1 {
		delete(this.FrequencyStack, this.MaxFreq)
		this.MaxFreq--
	} else {
		this.FrequencyStack[this.MaxFreq] = this.FrequencyStack[this.MaxFreq][:l-1]
	}

	newFreq := this.FrequencyMap[e] - 1
	if newFreq == 0 {
		delete(this.FrequencyMap, e)
	} else {
		this.FrequencyMap[e] = newFreq
	}

	return e
}

/* Notes:

-- Problem Link:
https://leetcode.com/problems/maximum-frequency-stack?envType=problem-list-v2&envId=7i18lqc

-- Notes:

1. Since we need the order the elements were pushed in to
resolve the tie, it is important to NOT update the
Frequency Stack of old frequency whenever the frequency
of an element increases. This will ensure that the older
frequency's frequency stack always has the order of the
elements preserved.


-- References:
https://rishabh.io/maximum-frequency-stack-e1e558e66201

*/
