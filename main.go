package main

import (
	"fmt"
	"github.com/sebzz2k2/goForDSA/array"
)

func main() {
	a := array.Array{
		Arr: []int{10, 21, 23, 87},
		Max: 5,
	}
	a.Insert(1, 89)
	a.Delete(2)
	a.Display()
	fmt.Println(a.Find(87))
	a.Reverse()
	fmt.Println(a.Arr)
}

	

