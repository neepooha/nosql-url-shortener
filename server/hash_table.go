package main

import "fmt"

const MAP_SIZE = 255

type HashMap struct {
	Data []*HashNode
}

func Constructor_HashMap() *HashMap {
	return &HashMap{Data: make([]*HashNode, MAP_SIZE)}
}

func (h *HashMap) HSET(key string, value string) {
	index := getIndex(key)

	if h.Data[index] == nil {
		h.Data[index] = &HashNode{key: key, value: value}
	} else {
		current := h.Data[index]
		for ; ; current = current.next {
			if current.key == key {
				// the key exists, its a modifying operation
				current.value = value
				return
			}
			if current.next == nil {
				break
			}
		}
		current.next = &HashNode{key: key, value: value}
	}
}

func (h *HashMap) HDEL(key string) bool {
	index := getIndex(key)
	current := h.Data[index]
	if current != nil {
		if current.key == key {
			h.Data[index] = current.next
			return true
		}
		for ; current.next != nil; current = current.next {
			if current.next.key == key {
				current.next = current.next.next
				return true
			}
		}
	}
	return false
}

func (h *HashMap) HGET(key string) (string, bool) {
	index := getIndex(key)
	if h.Data[index] != nil {
		current := h.Data[index]
		for ; ; current = current.next {
			if current.key == key {
				return current.value, true
			}

			if current.next == nil {
				break
			}
		}
	}

	// key does not exists
	return "", false
}

func hash(key string) (hash uint8) {
	// a jenkins one-at-a-time-hash
	// refer https://en.wikipedia.org/wiki/Jenkins_hash_function

	hash = 0
	for _, ch := range key {
		hash += uint8(ch)
		hash += hash << 10
		hash ^= hash >> 6
	}

	hash += hash << 3
	hash ^= hash >> 11
	hash += hash << 15

	return
}

func getIndex(key string) (index int) {
	return int(hash(key)) % MAP_SIZE
}

func (h HashMap) Print() {
	for i, val := range h.Data {
		fmt.Print("[", i, "]\t")
		for ; val != nil; val = val.next {
			fmt.Print(val.key, "_", val.value, " -> ")
		}
		fmt.Println("nil")
	}
}

func testHash() {
	hash := Constructor_HashMap()
	hash.HSET("1", "v1")
	hash.HSET("2", "v2")
	hash.HSET("1", "v3")
	hash.HSET("3", "v3")
	hash.Print()
	test, bo := hash.HGET("3")
	fmt.Println(test, " ", bo)

	hash.HDEL("3")

	test2, bo2 := hash.HGET("3")
	fmt.Println(test2, " ", bo2)

	hash.Print()
}
