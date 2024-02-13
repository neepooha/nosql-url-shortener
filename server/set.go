package main

import "fmt"

type Set struct {
	data map[string]struct{}
}

func Constructor_Set() *Set {
	return &Set{
		data: make(map[string]struct{}),
	}
}

func (s *Set) SADD(val string) {
	s.data[val] = struct{}{}
}

func (s *Set) SREM(val string) {
	if len(s.data) == 0 {
		return
	}
	delete(s.data, val)
}

func (s *Set) SISMEMBER(val string) bool {
	_, ok := s.data[val]
	return ok
}

func (s *Set) Print() {
	for key, _ := range s.data {
		fmt.Print(key, " ")
	}
	fmt.Println()
}

func testset() {
	set := Constructor_Set()
	set.SADD("2")
	set.SADD("1")
	set.SADD("3")
	set.SADD("5")
	set.SADD("7")
	set.SADD("3")
	set.Print()
	set.SREM("1")
	set.Print()
	fmt.Println(set.SISMEMBER("5"))
}
