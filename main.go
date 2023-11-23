package main

import (
	"fmt"
)

type Array struct {
	max int
	arr []int
}

func (re *Array) insert(pos int, num int) {
	if pos < 0 || pos > len(re.arr) || !(len(re.arr)+1<=re.max) {
		fmt.Println("Position out of range")
		return
	}
	re.arr = append(re.arr, 0)
	copy(re.arr[pos+1:], re.arr[pos:])
	re.arr[pos] = num
}


func main() {
	a := Array{
		max: 5,
		arr: []int{10,21,23,89},
	}
	// a.insert(0, 10)
	// a.insert(1,21)
	// a.insert(2,23)
	// a.insert(1,89)

	fmt.Println(a.arr)
}
