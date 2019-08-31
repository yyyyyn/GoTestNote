package queuelist

import (
	"container/list"
	"fmt"
)

func StackTest() {
	stack := list.New()
	for i := 0; i < 5; i++ {
		stack.PushBack(i) // 入栈
	}
	for {
		if stack.Len() == 0 {
			break
		}
		ele := stack.Back() //取出链表最后一个元素
		val := ele.Value.(int)
		fmt.Printf("val: %d", val)
		stack.Remove(ele) //将最后一个元素删除，出栈
	}
}

func QueueTest() {
	queue := list.New()
	for i := 0; i < 5; i++ {
		queue.PushBack(i) // 入队列
	}
	for {
		if queue.Len() == 0 {
			break
		}
		ele := queue.Front() // 取出链表第一个元素
		val := ele.Value.(int)
		fmt.Printf("val: %d", val)
		queue.Remove(ele) // 将第一个元素删除，出队列
	}
}
