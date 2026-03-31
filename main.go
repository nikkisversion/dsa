package main

import (
	"fmt"

	streamofcharacters "github.com/nikkisversion/dsa/leetcode/system_design/2_data_stream_processing/2_stream_of_characters"
)

func main() {
	stream := streamofcharacters.Constructor([]string{"ab", "ba", "aaab", "abab", "baa"})

	fmt.Println(stream.Query(byte('a')))
	fmt.Println(stream.Query(byte('a')))
	fmt.Println(stream.Query(byte('a')))
	fmt.Println(stream.Query(byte('a')))
	fmt.Println(stream.Query(byte('a')))
	fmt.Println(stream.Query(byte('b')))
	fmt.Println(stream.Query(byte('a')))
	fmt.Println(stream.Query(byte('b')))
	fmt.Println(stream.Query(byte('a')))
	fmt.Println(stream.Query(byte('b')))
	fmt.Println(stream.Query(byte('b')))
	fmt.Println(stream.Query(byte('b')))
}
