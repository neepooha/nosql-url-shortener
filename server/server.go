package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
)

const (
	HOST = "localhost"
	PORT = "6378"
	TYPE = "tcp"
)

type data_stuct struct {
	set   *Set
	stack *Stack
	queue *Queue
	hmap  *HashMap
	mutex sync.Mutex
}

func Constructor_data() *data_stuct {
	return &data_stuct{
		set:   Constructor_Set(),
		stack: Constructor_Stack(),
		queue: Constructor_Queue(),
		hmap:  Constructor_HashMap(),
	}
}

func main() {
	listen, err := net.Listen(TYPE, HOST+":"+PORT)
	if err != nil {
		log.Fatalf("Port listening error: %s", err)
	}
	defer listen.Close()
	fmt.Println("The server is running. Waiting for connections...")
	data := Constructor_data()
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatalf("Connection error: %s", err)
		}
		go data.handleConnection(conn)
	}
}

func (data *data_stuct) handleConnection(conn net.Conn) {
	r := bufio.NewReader(conn)
	w := bufio.NewWriter(conn)

	key, _ := r.ReadString('\n')
	key = strings.TrimSpace(key)

	val, _ := r.ReadString('\n')
	val = strings.TrimSpace(val)

	fmt.Println("Get message: ", key, " ", val)
	data.mutex.Lock()
	if val == "-1" {
		value, ok := data.hmap.HGET(key)
		if ok {
			fmt.Fprint(w, value+"\n")
		} else {
			fmt.Fprint(w, "404\n")
		}
	} else {
		_, ok := data.hmap.HGET(key)
		if !ok {
			data.hmap.HSET(key, val)
			fmt.Fprint(w, "ok\n")
		}
	}
	data.mutex.Unlock()
	w.Flush()
	conn.Close()
}
