package main

import "fmt"

type Queue struct {
	head *Node
	tail *Node
}

func Constructor_Queue() *Queue {
	return &Queue{
		head: nil,
		tail: nil,
	}
}

func (q *Queue) QPUSH(str string) {
	node := Node_Constructor(str, nil)
	if q.tail == nil {
		q.head = node
		q.tail = node
	} else {
		q.tail.next = node
		q.tail = node
	}
}

func (q *Queue) QPOP() (string, bool) {
	if q.head == nil {
		return "", false
	} else {
		data := q.head.data
		q.head = q.head.next
		if q.head == nil {
			q.tail = nil
		}
		return data, true
	}
}

func (q Queue) Print() {
	fmt.Print("head -> ")
	for q.head != nil {
		fmt.Print(q.head.data, " -> ")
		q.head = q.head.next
	}
	fmt.Println("tail")
}

func testQeue() {
	que := Constructor_Queue()
	que.QPUSH("1")
	que.QPUSH("2")
	que.QPUSH("3")
	que.Print()
	que.QPOP()
	que.QPOP()
	_, t1 := que.QPOP()
	_, t2 := que.QPOP()
	que.Print()
	fmt.Println(t1, " ", t2)
}
