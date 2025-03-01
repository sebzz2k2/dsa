package lists

import (
	"errors"
	"fmt"
)

type QNode struct {
	next *QNode
	val  int
}

type Queue struct {
	head *QNode
	tail *QNode
	size int
}

func QConstructor() *Queue {
	return &Queue{head: nil, tail: nil, size: 0}
}
func (q *Queue) GetLength() int {
	return q.size
}

func (q *Queue) Push(val int) {
	newNode := &QNode{val: val}
	if q.tail != nil {
		q.tail.next = newNode
	} else {
		q.head = newNode
	}
	q.tail = newNode
	q.size += 1
}

func (q *Queue) Pop() (int, error) {
	if q.head == nil {
		return 0, errors.New("Queue is empty")
	}
	n := q.head
	q.head = n.next
	if q.head == nil {
		q.tail = nil
	}
	q.size -= 1
	return n.val, nil
}

func (q *Queue) Display() {
	if q.head == nil {
		fmt.Println("Em")
	}
	curr := q.head
	fmt.Println("")
	for curr != nil {
		if curr.next != nil {
			fmt.Print(curr.val, " -> ")
		} else {
			fmt.Print(curr.val, " -> nil")
		}
		curr = curr.next
	}
}
