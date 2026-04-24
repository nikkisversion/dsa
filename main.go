package main

import (
	"fmt"

	insertdeletegetrandomduplicates "github.com/nikkisversion/dsa/leetcode/system_design/3_data_structure_design/2_insert_delete_get_random_duplicates"
)

func main() {
	e := insertdeletegetrandomduplicates.Constructor()

	fmt.Println(e.Insert(4))
	fmt.Println(e.Insert(3))
	fmt.Println(e.Insert(4))
	fmt.Println(e.Insert(2))
	fmt.Println(e.Insert(4))

	fmt.Println(e.Remove(4))
	fmt.Println(e.Remove(3))
	fmt.Println(e.Remove(4))
	fmt.Println(e.Remove(4))

}
