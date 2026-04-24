package allooneds

import "slices"

type AllOne struct {
	KeyToCount  map[string]int
	CountToKeys map[int][]string
	MaxCount    int
	MinCount    int
}

func Constructor() AllOne {
	return AllOne{
		KeyToCount:  make(map[string]int),
		CountToKeys: make(map[int][]string),
		MaxCount:    0,
		MinCount:    0,
	}
}

func (this *AllOne) Inc(key string) {
	v, ok := this.KeyToCount[key]
	if !ok {
		this.KeyToCount[key] = 1
		this.CountToKeys[1] = append(this.CountToKeys[1], key)
	} else {
		v++
		this.KeyToCount[key] = v
		this.CountToKeys[v] = append(this.CountToKeys[v], key)
	}

	if v >= this.MaxCount {
		this.MaxCount = v
	}

	if v <= this.MinCount {
		this.MinCount = v
	}
}

func (this *AllOne) Dec(key string) {
	v, _ := this.KeyToCount[key]

	v = v - 1

	if v == 0 {
		delete(this.KeyToCount, key)
		if len(this.CountToKeys[1]) == 1 {
			delete(this.CountToKeys, 1)
		} else {
			index := slices.Index(this.CountToKeys[1], key)
			this.CountToKeys[1] = slices.Delete(this.CountToKeys[1], index, index+1)
		}
	} else {
		this.KeyToCount[key] = v
		this.CountToKeys[v] = append(this.CountToKeys[v], key)
	}

	if v >= this.MaxCount {
		this.MaxCount = v
	}

	if v <= this.MinCount {
		this.MinCount = v
	}
}

func (this *AllOne) GetMaxKey() string {

	arr := this.CountToKeys[this.MaxCount]
	if len(arr) == 0 {
		return ""
	}
	return arr[0]

}

func (this *AllOne) GetMinKey() string {

	arr := this.CountToKeys[this.MinCount]
	if len(arr) == 0 {
		return ""
	}
	return arr[0]

}
