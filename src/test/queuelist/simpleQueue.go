package queuelist

import (
	"fmt"
	"time"
)

type node struct {
	value interface{}
	prev  *node
	next  *node
}

type LinkedQueue struct {
	head *node
	tail *node
	size int
}

func (queue *LinkedQueue) Size() int {
	return queue.size
}

func (queue *LinkedQueue) Peek() interface{} {
	if queue.head == nil {
		//panic("Empty queue.")
		fmt.Println("empty queue")
		return ""
	}
	return queue.head.value
}

func (queue *LinkedQueue) Add(value interface{}) {
	new_node := &node{value, queue.tail, nil}
	if queue.tail == nil {
		queue.head = new_node
		queue.tail = new_node
	} else {
		queue.tail.next = new_node
		queue.tail = new_node
	}
	queue.size++
	new_node = nil
}

func (queue *LinkedQueue) Remove() {
	if queue.head == nil {
		panic("Empty queue.")
	}
	first_node := queue.head
	queue.head = first_node.next
	first_node.next = nil
	first_node.value = nil
	queue.size--
	first_node = nil
}

func TestSimplrQueue() {
	var jobs = new(LinkedQueue)
	go func() {
		for num := 0; num < 10; num++ {
			fmt.Println("To add node...")
			jobs.Add(time.Now())
			time.Sleep(time.Second * 1)
		}
	}()
	go func() {
		for {
			fmt.Println("To read node...")
			if jobs.Size() > 0 {
				fmt.Println("time is ", jobs.Peek())
				jobs.Remove()
			}
			time.Sleep(time.Second * 2)
		}
	}()
}
