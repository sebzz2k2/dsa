package main

import (
	// "github.com/sebzz2k2/goForDSA/array"
	"fmt"

	"github.com/sebzz2k2/goForDSA/sort"
)

func main() {
	arr := [10]int{32, 23, 54, 553, 534, 2432, 54, 12, 98, 5}
	sorted := sort.SelectionSort(arr[:])
	fmt.Println(sorted)

	// fmt.Println(array.MergeAlternately("abcd", "pq"))
}
