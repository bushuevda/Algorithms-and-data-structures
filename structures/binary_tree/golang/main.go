package main

import (
	"bst/bst"
	"fmt"
)

func main() {
	var comparator bst.Comparator[int] = func(a, b int) int {
		if a > b {
			return 1
		} else if a < b {
			return -1
		} else {
			return 0
		}
	}

	bt := &bst.BST[int, string]{Key: 3, ComparatorKeys: comparator}
	bt.Insert(2, "s")

	bt.Insert(1, "e")

	bt.PreOrderRecursive(func(val int) { fmt.Println(val) })

	bt = bt.RemoveNode(3, comparator)
	fmt.Println(bt)
	bt.PreOrderRecursive(func(val int) { fmt.Println(val) })

}
