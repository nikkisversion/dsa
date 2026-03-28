package maxfreqstack

import "testing"

func TestFreqStack_Push(t *testing.T) {
	type fields struct {
		FrequencyMap   map[int]int
		FrequencyStack map[int][]int
		MaxFreq        int
	}
	type args struct {
		val int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &FreqStack{
				FrequencyMap:   tt.fields.FrequencyMap,
				FrequencyStack: tt.fields.FrequencyStack,
				MaxFreq:        tt.fields.MaxFreq,
			}
			this.Push(tt.args.val)
		})
	}
}
