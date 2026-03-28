package main

import (
	"fmt"

	maxfreqstack "github.com/nikkisversion/dsa/leetcode/system_design/3_Max_Freq_Stack"
)

func main() {
	mfs := maxfreqstack.Constructor()

	mfs.Push(5)
	mfs.Push(7)
	mfs.Push(5)
	mfs.Push(7)
	mfs.Push(4)
	mfs.Push(5)
	fmt.Println(mfs.Pop())
	fmt.Println(mfs.Pop())
	fmt.Println(mfs.Pop())
	fmt.Println(mfs.Pop())
}
