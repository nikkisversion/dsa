package streamofcharacters

import (
	"strings"
)

type Node struct {
	isLastLetterOfWord bool
	Children           map[string]*Node
}

func NewNode(data string) *Node {
	return &Node{isLastLetterOfWord: false, Children: make(map[string]*Node)}
}

type Trie struct {
	RootNode *Node
}

func NormalizeString(s string) string {
	return strings.ToLower(strings.ReplaceAll(s, " ", ""))
}

func NewTrie() *Trie {
	return &Trie{RootNode: NewNode("")}
}

func (this *Trie) Insert(newString string) {
	str := NormalizeString(newString)

	currentNode := this.RootNode

	l := len(str)
	for i := l - 1; i >= 0; i-- {
		strE := string(newString[i])
		nextNode, ok := currentNode.Children[strE]
		if !ok {
			currentNode.Children[strE] = NewNode(strE)
			currentNode = currentNode.Children[strE]
			continue
		}
		currentNode = nextNode
	}
	currentNode.isLastLetterOfWord = true
}

func (this *Trie) Search(str string) bool {

	str = NormalizeString(str)

	currentNode := this.RootNode
	for _, e := range str {
		strE := string(e)

		nextNode, ok := currentNode.Children[strE]
		if !ok {
			return false
		}

		currentNode = nextNode
	}

	if currentNode.isLastLetterOfWord {
		return true
	}
	return false
}

type StreamChecker struct {
	Words            []string
	Storage          *Trie
	PossibleSuffixes string
	MaxLen           int
}

func Constructor(words []string) StreamChecker {

	ml := 0
	newTrie := NewTrie()

	for _, word := range words {
		if len(word) > ml {
			ml = len(word)
		}
		newTrie.Insert(word)
	}

	return StreamChecker{
		Words:            words,
		Storage:          newTrie,
		PossibleSuffixes: "",
		MaxLen:           ml,
	}

}

func (this *StreamChecker) Query(letter byte) bool {

	this.PossibleSuffixes = string(letter) + this.PossibleSuffixes
	if len(this.PossibleSuffixes) > this.MaxLen {
		this.PossibleSuffixes = this.PossibleSuffixes[:this.MaxLen]
	}

	if len(this.Storage.RootNode.Children) == 0 {
		return false
	}

	currentNode := this.Storage.RootNode
	for i := 0; i < len(this.PossibleSuffixes); i++ {

		ch := this.PossibleSuffixes[i]

		val, ok := currentNode.Children[string(ch)]
		if !ok {
			return false
		}

		if val.isLastLetterOfWord {
			return true
		}

		currentNode = val

	}

	return false

}

/*
Notes

-- Comments:
1. We use a trie to optmimise on time
2. The strings are inserted in the trie in
reverse order and the input stream is formed
in reverse order to account for suffix checking
foremost
3. We used recursion in one solution but that added a lot of
overhead
4. A key point was to not traverse the whole input stream
everytime a new input arrives. Because for n inputs, it will
require traversing an n length input stream and then n
iterations to check. n square time complexity. So we use
the MaxLen variable.

-- Problem Link:
https://leetcode.com/problems/stream-of-characters?envType=problem-list-v2&envId=ssd-ssd2-data-stream-processing


-- References:
1. https://medium.com/@itachisasuke/implementing-a-search-engine-in-golang-trie-data-structure-c45152ddda24
2. https://gist.github.com/pjcalvo/59563a2024e5d4966b6c28afe6d8fdc2


*/
