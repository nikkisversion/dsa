package streamofcharacters

import (
	"slices"
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
	// for _, e := range str {
	// 	strE := string(e)
	// 	nextNode, ok := currentNode.Children[strE]
	// 	if !ok {
	// 		currentNode.Children[strE] = NewNode(strE)
	// 		currentNode = currentNode.Children[strE]
	// 		continue
	// 	}
	// 	currentNode = nextNode
	// }
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

	newTrie := NewTrie()
	//reveresedWordsArray := []string{}
	for _, word := range words {
		//reverseStr := reverseString(word)
		//reveresedWordsArray = append(reveresedWordsArray, word)
		newTrie.Insert(word)
	}
	return StreamChecker{
		Words:            words,
		Storage:          newTrie,
		PossibleSuffixes: "",
	}

}

func reverseString(str string) string {
	runes := []rune(str)
	slices.Reverse(runes)
	return string(runes)
}

func (this *StreamChecker) Query(letter byte) bool {

	this.PossibleSuffixes = string(letter) + this.PossibleSuffixes

	if len(this.Storage.RootNode.Children) == 0 {
		return false
	}

	return this.query(0, this.Storage.RootNode)

}

func (this *StreamChecker) query(index int, currentNode *Node) bool {

	if index >= len(this.PossibleSuffixes) {
		return false
	}

	firstLetter := this.PossibleSuffixes[index]

	val, ok := currentNode.Children[string(firstLetter)]
	if ok {
		if val.isLastLetterOfWord {
			return true
		}
		return this.query(index+1, val)
	}

	return false

}

func (this *StreamChecker) Query1(letter byte) bool {
	this.PossibleSuffixes = string(letter) + this.PossibleSuffixes

	strToCheck := ""
	for i := 0; i < len(this.PossibleSuffixes); i++ {
		strToCheck = strToCheck + string(this.PossibleSuffixes[i])

		isExists := this.Storage.Search(strToCheck)
		if isExists {
			return true
		}

	}
	return false
}

func (this *StreamChecker) Query2(letter byte) bool {

	this.PossibleSuffixes = string(letter) + this.PossibleSuffixes

	// firstLetter := string(letter)

	// exists := this.Storage.Search(firstLetter)
	// if !exists {
	// 	return false
	// }

	// return true

	if !this.Storage.Search(string(this.PossibleSuffixes[0])) {
		return false
	}

	firstLetter := string(this.PossibleSuffixes[0])
	currentNode := this.Storage.RootNode.Children[firstLetter]

	for i := 1; i < len(this.PossibleSuffixes); i++ {

		letter := string(this.PossibleSuffixes[i])

		if len(currentNode.Children) == 0 {
			return true
		}

		val, ok := currentNode.Children[letter]
		if !ok {
			return false
		}

		currentNode = val

	}

	if len(currentNode.Children) == 0 {
		return true
	}

	return false

}
