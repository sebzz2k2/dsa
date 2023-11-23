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

func (re *Array) delete(pos int){
	if pos < 1 || pos > len(re.arr) {
		fmt.Println("Position out of range")
		return
	}
	re.arr= append(re.arr[:pos] , re.arr[pos+1:]...)
}

func main() {
	a := Array{
		max: 5,
		arr: []int{10,21,23,89,87},
	}
	a.insert(1,89)
	a.delete(2)

	fmt.Println(a.arr)
}
