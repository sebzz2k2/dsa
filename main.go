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

func (re *Array) display(){
	fmt.Println(re.arr)
}

func (re *Array) find(num int) int{
	for i := 0; i < len(re.arr); i++ {
		if re.arr[i]==num {
			return i
		}		
	}
	return -1
}

func (re *Array) reverse() {
    n := len(re.arr)
    for i := 0; i < n/2; i++ {
        j := n - i - 1
        if i != j {
			// bit manupulation for fun
            re.arr[i] = re.arr[i] ^ re.arr[j]
            re.arr[j] = re.arr[i] ^ re.arr[j]
            re.arr[i] = re.arr[i] ^ re.arr[j]
        }
    }
}


func main() {
	a := Array{
		max: 5,
		arr: []int{10,21,23,87},
	}
	a.insert(1,89)
	a.delete(2)
	a.display()
	fmt.Println(a.find(87))
	a.reverse()
	fmt.Println(a.arr)
}
