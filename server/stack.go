package main

import "fmt"

type Stack struct {
	head *Node
}

func Constructor_Stack() *Stack {
	return &Stack{head: nil}
}

func (s *Stack) SPUSH(str string) {
	node := Node_Constructor(str, nil)
	if s.head == nil {
		s.head = node
	} else {
		node.next = s.head
		s.head = node
	}
}

func (s *Stack) SPOP() (string, bool) {
	if s.head == nil {
		return "", false
	} else {
		element := s.head.data
		s.head = s.head.next
		return element, true
	}
}

func (s Stack) Print() {
	fmt.Print("head -> ")
	for ; s.head != nil; s.head = s.head.next {
		fmt.Print(s.head.data, " -> ")
	}
	fmt.Println("nil")
}

func testStack() {
	sta := Constructor_Stack()
	sta.SPUSH("1")
	sta.SPUSH("2")
	sta.SPUSH("3")
	sta.Print()
	sta.SPOP()
	sta.SPOP()
	_, t1 := sta.SPOP()
	_, t2 := sta.SPOP()
	sta.Print()
	fmt.Println(t1, " ", t2)
}
