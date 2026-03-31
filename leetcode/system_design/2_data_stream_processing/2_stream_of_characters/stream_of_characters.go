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
